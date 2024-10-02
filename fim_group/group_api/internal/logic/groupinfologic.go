package logic

import (
	"context"
	"encoding/json"

	"fim/fim_group/group_api/internal/svc"
	"fim/fim_group/group_api/internal/types"
	"fim/fim_group/group_rpc/types/group_rpc"
	groupmodel "fim/fim_group/models"
	usermodel "fim/fim_user/models"
	"fim/fim_user/user_rpc/types/user_grpc"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupInfoLogic {
	return &GroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupInfoLogic) GroupInfo(req *types.GroupInfoRequest) (resp *types.GroupInfoResponse, err error) {
	// todo: add your logic here and delete this line
	//必须该群成员才可以调接口

	//判断群是否存在
	var group groupmodel.GroupModel
	err = l.svcCtx.DB.Where("id = ?", req.GroupId).First(&group).Error
	if err != nil {
		logx.Error("l.svcCtx.DB.Where error", err)
		return nil, errors.New("Group doesn't exist")
	}

	//判断是否是群成员
	var groupMember groupmodel.GroupMembersModel
	err = l.svcCtx.DB.Where("group_id = ? and user_id = ?", req.GroupId, req.UserId).First(&groupMember).Error
	if err != nil {
		logx.Error("l.svcCtx.DB.Where error", err)
		return nil, errors.New("You are not a member of this group")
	}
	//查看群成员信息
	var groupMembers []groupmodel.GroupMembersModel
	err = l.svcCtx.DB.Where("group_id = ?", req.GroupId).Preload("UserModel").Find(&groupMembers).Error
	if err != nil {
		logx.Error("l.svcCtx.DB.Where error", err)
		return nil, errors.New("Group members query failed")
	}
	for _, v := range groupMembers {
		_, err := l.svcCtx.GroupRpc.GroupMemberInfo(context.Background(), &group_rpc.MemberInfoRequest{
			Group:  uint32(req.GroupId),
			Member: uint32(v.UserID),
		})
		if err != nil {
			logx.Error("l.svcCtx.GroupRpc.GroupMemberInfo error", err)
			return nil, errors.New("Group members query failed")
		}

	}
	//查在线人数
	var userOnline []usermodel.UserConfigModel
	for _, v := range groupMembers {
		var userConf usermodel.UserConfigModel
		err := l.svcCtx.DB.Where("user_id = ?", v.UserID).First(&userConf).Error
		if err != nil {
			logx.Error("l.svcCtx.DB.Where error", err)
			return nil, errors.New("UserConf doesn't exist")
		}
		if userConf.Online == true {
			userOnline = append(userOnline, userConf)
		}
	}
	//查群主信息
	var creator usermodel.UserModel
	res, err := l.svcCtx.UserRpc.UserInfo(context.Background(), &user_grpc.UserInfoRequest{
		UserId: uint32(group.CreatorID),
	})
	if err != nil {
		logx.Error("l.svcCtx.UserRpc.UserInfo error", err)
		return nil, errors.New("Creator doesn't exist")
	}
	err = json.Unmarshal(res.Data, &creator)
	if err != nil {
		logx.Error("json.Unmarshal error", err)
		return nil, errors.New("json.Unmarshal error")
	}

	//查管理员信息
	var adminList []types.UserInfo
	for _, v := range groupMembers {
		if v.Role != 0 {
			admin := types.UserInfo{
				UserId:   v.UserID,
				UserName: v.MemberNickName,
				Avatar:   v.UserModel.Avatar,
			}
			adminList = append(adminList, admin)
		}
	}

	return &types.GroupInfoResponse{
		GroupId:     group.ID,
		Title:       group.Title,
		Abstract:    group.Abstract,
		Avatar:      group.Avatar,
		MemberCount: (len(groupMembers)),
		OnlineCount: len(userOnline), //获取到了groupMembers之后，查这些人是否在线
		Creator: types.UserInfo{
			UserId:   creator.ID,
			UserName: creator.NickName,
			Avatar:   creator.Avatar,
		},
		AdminList: adminList, //查Role不是0的人，返回
	}, nil
}

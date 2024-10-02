package logic

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"fim/fim_group/group_api/internal/svc"
	"fim/fim_group/group_api/internal/types"
	groupmodel "fim/fim_group/models"
	usermodel "fim/fim_user/models"
	"fim/fim_user/user_rpc/types/user_grpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCreateLogic {
	return &GroupCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupCreateLogic) GroupCreate(req *types.GroupCreateRequest) (resp *types.GroupCreateResponse, err error) {
	// todo: add your logic here and delete this line
	var group = groupmodel.GroupModel{
		CreatorID:    req.UserId,
		IsSearchable: false,
		Verification: 2,
		Size:         1,
	}

	switch req.Mode {
	case 1: //直接创建
		if req.Name == "" {
			return nil, errors.New("Group name is empty")
		}
		group.Title = req.Name
		group.Size = int8(req.Size)
		group.IsSearchable = req.IsSearch
		group.Avatar = "fim_file/file_api/uploads/avatar/IMG_7525.jpeg"

	case 2: //选人创建
		if len(req.UserIdList) == 0 {
			return nil, errors.New("UserIdList is empty")
		}
		var userIdList = []uint32{
			uint32(req.UserId), //先把自己加进去
		}
		var nameList []string
		for _, v := range req.UserIdList {
			userIdList = append(userIdList, uint32(v))
			res, err := l.svcCtx.UserRpc.UserInfo(context.Background(), &user_grpc.UserInfoRequest{
				UserId: uint32(v),
			})
			if err != nil {
				return nil, err
			}

			user := usermodel.UserModel{}
			err = json.Unmarshal(res.Data, &user)
			if err != nil {
				return nil, errors.New("json.Unmarshal(res.Data,&user) error")
			}

			if len(strings.Join(nameList, "、")) >= 32 {
				break
			}
			nameList = append(nameList, user.NickName)
		}
		nameList[len(nameList)-1] = nameList[len(nameList)-1] + "的群聊"
		group.Title = strings.Join(nameList, "、")
		group.Avatar = "fim_file/file_api/uploads/avatar/IMG_7525.jpeg"
		group.IsSearchable = req.IsSearch
	}

	err = l.svcCtx.DB.Create(&group).Error
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	if req.UserIdList != nil {
		for _, v := range req.UserIdList {
			res, err := l.svcCtx.UserRpc.UserInfo(context.Background(), &user_grpc.UserInfoRequest{
				UserId: uint32(v),
			})
			if err != nil {
				logx.Error(err)
				return nil, err
			}
			var user = usermodel.UserModel{}
			err = json.Unmarshal(res.Data, &user)
			var groupMember = groupmodel.GroupMembersModel{
				GroupID:        group.ID,
				GroupModel:     group,
				UserID:         uint(v),
				UserModel:      user,
				MemberNickName: user.NickName,
				Role:           0,
				IsMute:         false,
			}
			err = l.svcCtx.DB.Create(&groupMember).Error
			if err != nil {
				logx.Error(err)
				return nil, err
			}
		}
	}

	return &types.GroupCreateResponse{
		GroupId: group.ID,
	}, nil
}

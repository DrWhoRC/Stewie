package logic

import (
	"context"
	"errors"

	"fim/fim_group/group_rpc/internal/svc"
	"fim/fim_group/group_rpc/types/group_rpc"
	groupmodel "fim/fim_group/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMemberInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupMemberInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMemberInfoLogic {
	return &GroupMemberInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupMemberInfoLogic) GroupMemberInfo(in *group_rpc.MemberInfoRequest) (*group_rpc.MemberInfoResponse, error) {
	// todo: add your logic here and delete this line

	var groupMember groupmodel.GroupMembersModel
	err := l.svcCtx.DB.Where("group_id = ? and user_id = ?", in.Group, in.Member).Preload("UserModel").Find(&groupMember).Error
	if err != nil {
		logx.Error("l.svcCtx.DB.Where error", err)
		return nil, errors.New("Group member doesn't exist")
	}

	// var user usermodel.UserModel
	// err = l.svcCtx.DB.Where("id = ?", in.Member).Find(&user).Error
	// if err != nil {
	// 	logx.Error("l.svcCtx.DB.Where error", err)
	// 	return nil, errors.New("User doesn't exist")
	// }
	//上边用了Preload函数，这里就不用再查一遍user表，只要在model文件中写了外键进行关联了就行

	return &group_rpc.MemberInfoResponse{
		UserId:   int32(groupMember.UserID),
		UserName: groupMember.MemberNickName,
		Avatar:   groupMember.UserModel.Avatar,
		Role:     int32(groupMember.Role),
	}, nil
}

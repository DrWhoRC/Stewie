package logic

import (
	"context"
	"fmt"

	usermodel "fim/fim_user/models"
	"fim/fim_user/user_rpc/internal/svc"
	"fim/fim_user/user_rpc/types/user_grpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendListLogic) GetFriendList(in *user_grpc.FriendListRequest) (*user_grpc.FriendListResponse, error) {
	// todo: add your logic here and delete this line

	var users []usermodel.FriendModel
	err := l.svcCtx.DB.Model(&usermodel.FriendModel{}).
		Where("sender_id = ? OR receiver_id = ?", in.UserId, in.UserId).Find(&users).Error
	if err != nil {
		logx.Error("GetFriendListLogic GetFriendList SQL error", err)
		return &user_grpc.FriendListResponse{}, err
	}

	list := []*user_grpc.FriendInfo{}

	for _, user := range users {
		var id uint
		var notice string
		if user.SenderID == uint(in.UserId) {
			id = user.ReceiverID
			notice = user.RecvUserNotice
		} else {
			id = user.SenderID
			notice = user.SendUserNotice
		}

		var user usermodel.UserModel
		l.svcCtx.DB.Preload("UserConfigModel").Model(&usermodel.UserModel{}).Where("ID = ?", id).First(&user)
		list = append(list, &user_grpc.FriendInfo{
			UserId:             uint32(id),
			Nickname:           fmt.Sprintf("%s(%s)", user.NickName, notice),
			Avatar:             user.Avatar,
			FriendOnlineNotify: user.UserConfigModel.FriendOnlineNotify,
		})
	}

	return &user_grpc.FriendListResponse{
		FriendList: list,
	}, nil
}

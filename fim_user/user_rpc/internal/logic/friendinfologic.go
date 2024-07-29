package logic

import (
	"context"
	"encoding/json"

	usermodel "fim/fim_user/models"
	"fim/fim_user/user_rpc/internal/svc"
	"fim/fim_user/user_rpc/types/user_grpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendInfoLogic {
	return &FriendInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendInfoLogic) FriendInfo(in *user_grpc.FriendInfoRequest) (*user_grpc.FriendInfoResponse, error) {
	// todo: add your logic here and delete this line

	var user usermodel.FriendModel
	//var friend usermodel.FriendModel
	err := l.svcCtx.DB.Model(&usermodel.FriendModel{}).Where("ID = ?", uint(in.UserId)).First(&user).Error
	if err != nil {
		return &user_grpc.FriendInfoResponse{
			Data: []byte("user ID not found"),
		}, err
	}
	bytedata1, _ := json.Marshal(user.Notice)

	var friend usermodel.UserModel
	err = l.svcCtx.DB.Model(&usermodel.UserModel{}).Where("ID = ?", uint(in.FriendId)).First(&friend).Error
	if err != nil {
		return &user_grpc.FriendInfoResponse{
			Data: []byte("frined ID not found"),
		}, err
	}
	bytedata2, _ := json.Marshal(friend)
	bytedata := append(bytedata1, bytedata2...)
	return &user_grpc.FriendInfoResponse{
		Data: bytedata,
	}, nil
}

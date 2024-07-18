package logic

import (
	"context"
	usermodel "fim/fim_user/models"

	"fim/fim_user/user_rpc/internal/svc"
	"fim/fim_user/user_rpc/types/user_grpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user_grpc.UserInfoRequest) (*user_grpc.UserInfoResponse, error) {
	// todo: add your logic here and delete this line

	var user usermodel.UserModel
	err := l.svcCtx.DB.Where("id = ?", in.UserId).First(&user)
	if err != nil {
		return &user_grpc.UserInfoResponse{
			Data: []byte("user ID not found"),
		}, err.Error
	}

	return &user_grpc.UserInfoResponse{}, nil
}

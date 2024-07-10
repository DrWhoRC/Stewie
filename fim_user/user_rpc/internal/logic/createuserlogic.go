package logic

import (
	"context"

	"fim/fim_user/user_rpc/internal/svc"
	"fim/fim_user/user_rpc/types/user_grpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user_grpc.UserCreateRequest) (*user_grpc.UserCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &user_grpc.UserCreateResponse{}, nil
}

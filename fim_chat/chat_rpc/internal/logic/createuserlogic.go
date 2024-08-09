package logic

import (
	"context"

	"fim/fim_chat/chat_rpc/internal/svc"
	"fim/fim_chat/chat_rpc/types/chat_rpc"

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

func (l *CreateUserLogic) CreateUser(in *chat_rpc.UserChatRequest) (*chat_rpc.UserChatResponse, error) {
	// todo: add your logic here and delete this line

	return &chat_rpc.UserChatResponse{}, nil
}

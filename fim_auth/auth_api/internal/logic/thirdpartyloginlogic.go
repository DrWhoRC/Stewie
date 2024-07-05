package logic

import (
	"context"

	"fim/fim_auth/auth_api/internal/svc"
	"fim/fim_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Third_party_loginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThird_party_loginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Third_party_loginLogic {
	return &Third_party_loginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Third_party_loginLogic) Third_party_login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

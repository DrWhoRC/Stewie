package logic

import (
	"context"

	"fim/fim_user/user_api/internal/svc"
	"fim/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserConfUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserConfUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserConfUpdateLogic {
	return &UserConfUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserConfUpdateLogic) UserConfUpdate(req *types.UserConfUpdateRequest) (resp *types.UserConfUpdateResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

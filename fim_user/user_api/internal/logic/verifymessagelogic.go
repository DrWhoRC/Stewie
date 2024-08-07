package logic

import (
	"context"

	usermodel "fim/fim_user/models"
	"fim/fim_user/user_api/internal/svc"
	"fim/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyMessageLogic {
	return &VerifyMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyMessageLogic) VerifyMessage(req *types.VerifyMessageRequest) (resp *types.VerifyMessageResponse, err error) {
	// todo: add your logic here and delete this line
	var userConf usermodel.UserConfigModel
	l.svcCtx.DB.Model(&usermodel.UserConfigModel{}).Where("user_id = ?", req.ReceiverId).First(&userConf)

	if req.Agree == 1 {
		AddToFriend(req.SenderId, req.ReceiverId, l.svcCtx.DB)
		resp.Data = "you have become friends with him/her"
	}
	if req.Agree == 0 {
		resp.Data = "you have rejected the friend request"
	}

	return
}

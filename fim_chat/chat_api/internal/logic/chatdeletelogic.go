package logic

import (
	"context"

	"fim/fim_chat/chat_api/internal/svc"
	"fim/fim_chat/chat_api/internal/types"
	chatmodel "fim/fim_chat/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatDeleteLogic {
	return &ChatDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatDeleteLogic) ChatDelete(req *types.ChatDeleteRequest) (resp *types.ChatDeleteResponse, err error) {
	// todo: add your logic here and delete this line

	var deletedChat chatmodel.UserChatDeleteModel
	deletedChat.ChatId = req.ChatId
	deletedChat.UserId = req.UserId
	err = l.svcCtx.DB.Create(&deletedChat).Error
	if err != nil {
		logx.Error(err)
		return
	}

	return &types.ChatDeleteResponse{
		Data: "Chat Deleted Successfully",
	}, nil
}

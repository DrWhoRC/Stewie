package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"fim/common/models/ctype"
	"fim/fim_chat/chat_rpc/internal/svc"
	"fim/fim_chat/chat_rpc/types/chat_rpc"
	chatmodel "fim/fim_chat/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserChatLogic {
	return &UserChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserChatLogic) UserChat(in *chat_rpc.UserChatRequest) (*chat_rpc.UserChatResponse, error) {
	// todo: add your logic here and delete this line

	var msg ctype.Msg
	err := json.Unmarshal(in.Msg, &msg)
	if err != nil {
		logx.Error("json.Unmarshal(in.Msg, &msg) error", err)
		return nil, err
	}
	fmt.Println("req:", in.Sender, "+", in.Receiver, "+", in.Msg)

	err = l.svcCtx.DB.Create(&chatmodel.ChatModel{
		SenderID:   uint(in.Sender),
		ReceiverID: uint(in.Receiver),
		MsgType:    msg.Type,
		MsgPreview: MsgPreview(msg),
		Msg:        msg,
	}).Error
	if err != nil {
		logx.Error("l.svcCtx.DB.Create error", err)
		return nil, err
	}
	return &chat_rpc.UserChatResponse{
		UserId: int32(in.Sender),
	}, nil
}

func MsgPreview(msg ctype.Msg) string {
	switch msg.Type {
	case 0:
		return *msg.Content
	case 1:
		return "[picture]-" + msg.ImageMsg.Title
	case 2:
		return "[file]-" + msg.FileMsg.Title
	case 3:
		return "[voice]-"
	case 4:
		return "[video]-" + msg.VideoMsg.Title
	case 6:
		return "[voice call]-" + msg.VoiceCallMsg.StartTime.String()
	case 7:
		return "[video call]-" + msg.VideoCallMsg.StartTime.String()
	case 8:
		return "[Withdraw a message-]" + msg.WithdrawMsg.Content
	case 9:
		return "[Quote a message-]" + *msg.QuoteMsg.Msg.Content
	case 10:
		return "[@ Msg]-" + msg.AtMsg.Content
	default:
		return "you have a new message"
	}
}

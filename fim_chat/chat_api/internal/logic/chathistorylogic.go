package logic

import (
	"context"
	"fmt"

	"fim/common/models/ctype"
	"fim/fim_chat/chat_api/internal/svc"
	"fim/fim_chat/chat_api/internal/types"
	chatmodel "fim/fim_chat/models"
	usermodel "fim/fim_user/models"
	utils "fim/utils/dupicate"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type UserInfo struct {
	UserId   uint   `json:"userId"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}
type ChatHistory struct {
	MsgId     uint      `json:"msgId"`
	Sender    UserInfo  `json:"sender"`
	Receiver  UserInfo  `json:"receiver"`
	CreatedAt string    `json:"createdAt"`
	Msg       ctype.Msg `json:"msg"`
}

type ChatHistoryResponse struct {
	List  []ChatHistory `json:"list"`
	Count int           `json:"count"`
}

func NewChatHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatHistoryLogic {
	return &ChatHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatHistoryLogic) ChatHistory(req *types.ChatHistoryRequest) (resp *ChatHistoryResponse, err error) {
	// todo: add your logic here and delete this line

	if req.Limit < 0 {
		req.Limit = 10
	}
	if req.Page < 0 {
		req.Page = 1
	}
	offset := (req.Page - 1) * req.Limit

	var chatList []chatmodel.ChatModel
	l.svcCtx.DB.Model(&chatmodel.ChatModel{}).Limit(int(req.Limit)).Offset(int(offset)).Order("created_at DESC").Find(&chatList, "sender_id = ? or receiver_id = ?", req.UserId, req.UserId)

	var deletedChat []uint
	l.svcCtx.DB.Model(&chatmodel.UserChatDeleteModel{}).
		Select("chat_id").
		Where("user_id = ?", req.UserId).
		Find(&deletedChat)

	chatList = filterDeletedChats(chatList, deletedChat)

	var userIdList []int
	for _, model := range chatList {
		userIdList = append(userIdList, int(model.SenderID))
		userIdList = append(userIdList, int(model.ReceiverID))
	}
	//去重
	userIdList = utils.RemoveDuplicateElement(userIdList)

	fmt.Println(userIdList)

	var users []usermodel.UserModel
	l.svcCtx.DB.Model(&usermodel.UserModel{}).Where("id IN (?)", userIdList).Find(&users)

	// 创建 AvatarMap
	AvatarMap := make(map[int]string)
	for _, user := range users {
		AvatarMap[int(user.ID)] = user.Avatar
	}
	NicknameMap := make(map[int]string)
	for _, user := range users {
		NicknameMap[int(user.ID)] = user.NickName
	}

	var list []ChatHistory
	for _, v := range chatList {
		list = append(list, ChatHistory{
			MsgId: uint(v.ID),
			Sender: UserInfo{
				UserId:   uint(v.SenderID),
				Avatar:   AvatarMap[int(v.SenderID)],
				Nickname: NicknameMap[int(v.SenderID)],
			},
			Receiver: UserInfo{
				UserId:   uint(v.ReceiverID),
				Avatar:   AvatarMap[int(v.ReceiverID)],
				Nickname: NicknameMap[int(v.ReceiverID)],
			},
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			Msg:       v.Msg,
		})
	}

	resp = &ChatHistoryResponse{
		List:  list,
		Count: len(list),
	}
	return
}

func filterDeletedChats(chatList []chatmodel.ChatModel, deletedChat []uint) []chatmodel.ChatModel {
	deletedChatMap := make(map[uint]struct{}, len(deletedChat))
	for _, id := range deletedChat {
		deletedChatMap[id] = struct{}{}
	}

	var filteredList []chatmodel.ChatModel
	for _, chat := range chatList {
		if _, found := deletedChatMap[chat.ID]; !found {
			filteredList = append(filteredList, chat)
		}
	}
	return filteredList
}

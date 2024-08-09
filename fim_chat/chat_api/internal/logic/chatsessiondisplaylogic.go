package logic

import (
	"context"
	"fmt"
	"time"

	"fim/fim_chat/chat_api/internal/svc"
	"fim/fim_chat/chat_api/internal/types"
	chatmodel "fim/fim_chat/models"
	usermodel "fim/fim_user/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatSessionDisplayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatSessionDisplayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatSessionDisplayLogic {
	return &ChatSessionDisplayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 最近消息展示
func (l *ChatSessionDisplayLogic) ChatSessionDisplay(req *types.ChatSessionDisplayRequest) (resp *types.ChatSessionDisplayResponse, err error) {
	// todo: add your logic here and delete this line

	if req.Limit < 0 {
		req.Limit = 10
	}
	if req.Page < 0 {
		req.Page = 1
	}
	offset := (req.Page - 1) * req.Limit

	type Data struct {
		SU         uint      `gorm:"column:sU"`
		RU         uint      `gorm:"column:rU"`
		MaxDate    time.Time `gorm:"column:maxDate"`
		MaxPreview string    `gorm:"column:maxPreview"`
	}
	var dataList []Data

	l.svcCtx.DB.Table("(?) as u", l.svcCtx.DB.Model(&chatmodel.ChatModel{}).
		Select("least(sender_id, receiver_id)	as sU",
			"greatest(sender_id, receiver_id)	as rU",
			"max(created_at)	as maxDate",
			"max(msg_preview)	as maxPreview").
		//妈的咋改啊，捏妈妈的，又要join来join去确保这个msg-preview是最新的msg
		Where("sender_id = ? OR receiver_id = ?", req.UserId, req.UserId).
		Group("least(sender_id, receiver_id), greatest(sender_id, receiver_id)")).
		Order("maxDate desc").Limit(int(req.Limit)).Offset(int(offset)).
		Scan(&dataList)

	fmt.Println("dataList:", dataList[0])

	var userList []uint
	for _, v := range dataList {
		if v.RU != req.UserId {
			userList = append(userList, v.RU)
		}
		if v.SU != req.UserId {
			userList = append(userList, v.SU)
		}
	}
	fmt.Println(userList)

	var users []usermodel.UserModel
	l.svcCtx.DB.Model(&usermodel.UserModel{}).Where("id IN (?)", userList).Find(&users)

	AvatarMap := make(map[int]string)
	for _, user := range users {
		AvatarMap[int(user.ID)] = user.Avatar
	}
	NicknameMap := make(map[int]string)
	for _, user := range users {
		NicknameMap[int(user.ID)] = user.NickName
	}

	var list []types.ChatSessionDisplay
	for k, v := range dataList {
		list = append(list, types.ChatSessionDisplay{
			UserId:     userList[k],
			Avatar:     AvatarMap[int(userList[k])],
			Nickname:   NicknameMap[int(userList[k])],
			CreatedAt:  v.MaxDate.Format("2006-01-02 15:04:05"),
			MsgPreview: v.MaxPreview,
		})
	}

	resp = &types.ChatSessionDisplayResponse{
		List:  list,
		Count: len(list),
	}

	return
}

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
		IsTop      bool      `gorm:"column:isTop"`
	}
	var dataList []Data

	subQuery := l.svcCtx.DB.Model(&chatmodel.ChatModel{}).
		Select("least(sender_id, receiver_id) as sU",
			"greatest(sender_id, receiver_id) as rU",
			"max(created_at) as maxDate").
		Where("sender_id = ? OR receiver_id = ?", req.UserId, req.UserId).
		Group("least(sender_id, receiver_id), greatest(sender_id, receiver_id)")

	// 主查询：使用子查询获取最新消息的预览
	l.svcCtx.DB.Table("(?) as u", subQuery).
		Joins("JOIN chat_models cm ON cm.created_at = u.maxDate").
		Joins("LEFT JOIN top_user_models tu ON tu.user_id = ? AND (tu.top_user_id = u.sU OR tu.top_user_id = u.rU)", req.UserId).
		Select("u.sU, u.rU, u.maxDate, cm.msg_preview as maxPreview, CASE WHEN tu.id IS NOT NULL THEN 1 ELSE 0 END as isTop").
		Order("isTop DESC, u.maxDate desc").Limit(int(req.Limit)).Offset(int(offset)).
		Scan(&dataList)
	//使用 CASE 语句判断用户是否为置顶用户，并将结果存储在 IsTop 字段中。
	//在 ORDER BY 子句中，首先按 IsTop 字段降序排序，然后按 u.maxDate 字段降序排序。

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

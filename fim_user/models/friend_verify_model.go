package models

import (
	"fim/common/models/ctype"

	"gorm.io/gorm"
)

type FriendVerifyModel struct {
	gorm.Model
	SenderID       uint                  `json:"senderID"`
	ReceiverID     uint                  `json:"receiverID"`
	Status         int8                  `json:"status"`         //好友状态: 0-待确认; 1-已接受; 2-已拒绝
	VerifyQuestion *ctype.VerifyQuestion `json:"verifyQuestion"` //验证问题: Only needed when friendVerify=2 or 3
	Attached       string                `json:"attached"`       //附加信息
}

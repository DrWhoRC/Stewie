package models

import (
	"fim/common/models/ctype"

	"gorm.io/gorm"
)

type FriendVerifyModel struct {
	gorm.Model
	SenderID         uint                  `json:"senderID"`
	SendUserModel    UserModel             `gorm:"foreignKey:SenderID" json:"sendUserModel"`
	ReceiverID       uint                  `json:"receiverID"`
	ReceiveUserModel UserModel             `gorm:"foreignKey:ReceiverID" json:"receiveUserModel"`
	Status           int8                  `json:"status"`                   //好友状态: 0-待确认; 1-已接受; 2-已拒绝
	VerifyQuestion   *ctype.VerifyQuestion `json:"verifyQuestion"`           //验证问题: Only needed when friendVerify=2 or 3
	Attached         string                `gorm:"size:128" json:"attached"` //附加信息
}

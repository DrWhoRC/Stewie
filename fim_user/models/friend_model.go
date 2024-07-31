package models

import "gorm.io/gorm"

type FriendModel struct {
	gorm.Model
	SenderID         uint      `json:"senderID"`
	ReceiverID       uint      `json:"receiverID"`
	SendUserModel    UserModel `gorm:"foreignKey:SenderID" json:"sendUserModel"`
	ReceiveUserModel UserModel `gorm:"foreignKey:ReceiverID" json:"receiveUserModel"`
	SendUserNotice   string    `gorm:"size:128" json:"sendUserNotice"`
	RecvUserNotice   string    `gorm:"size:128" json:"recvUserNotice"`
	Status           int8      `json:"status"`                   //好友状态: 0-待确认; 1-已接受; 2-已拒绝
	Attached         string    `gorm:"size:128" json:"attached"` //附加信息
}

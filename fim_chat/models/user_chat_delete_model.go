package models

import (
	"gorm.io/gorm"
)

type UserChatDeleteModel struct {
	gorm.Model
	UserId uint `json:"userId"`
	ChatId uint `json:"chatId"`
}

package models

import (
	"gorm.io/gorm"
)

// pinned_user_model
type TopUserModel struct {
	gorm.Model
	UserId    uint `json:"userId"`
	TopUserId uint `json:"topUserId"`
}

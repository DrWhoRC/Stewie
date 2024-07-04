package models

import (
	"fim/common/models/ctype"
	usermodel "fim/fim_user/models"

	"gorm.io/gorm"
)

type GroupVerifyModel struct {
	gorm.Model
	GroupID        uint                 `json:"groupID"`
	GroupModel     GroupModel           `gorm:"foreignKey:GroupID" json:"groupModel"`
	UserID         uint                 `json:"userID"`
	UserModel      usermodel.UserModel  `gorm:"foreignKey:UserID" json:"userModel"`
	Attached       string               `gorm:"size:32" json:"attached"` //附加信息
	VerifyQuestion ctype.VerifyQuestion `json:"verifyQuestion"`
	Status         int8                 `json:"status"` //验证状态: 0-待确认; 1-已接受; 2-已拒绝
	Type           int8                 `json:"type"`   //验证类型: 0-申请加入; 1-退群
}

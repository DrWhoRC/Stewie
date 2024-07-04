package models

import (
	usermodel "fim/fim_user/models"

	"gorm.io/gorm"
)

type GroupMembersModel struct {
	gorm.Model
	GroupID        uint                `json:"groupID"`
	GroupModel     GroupModel          `gorm:"foreignKey:GroupID;references:ID" json:"groupModel"`
	UserID         uint                `json:"userID"`
	UserModel      usermodel.UserModel `gorm:"foreignKey:UserID;references:ID" json:"userModel"`
	MemberNickName string              `gorm:"size:32" json:"memberNickName"`
	Role           int8                `json:"role"`     //角色: 0-普通成员; 1-管理员; 2-群主
	IsMute         bool                `json:"isMute"`   //是否禁言
	MuteTime       int64               `json:"muteTime"` //禁言时间
}

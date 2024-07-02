package models

import "gorm.io/gorm"

type GroupMembersModel struct {
	gorm.Model
	GroupID        uint   `json:"groupID"`
	UserID         uint   `json:"userID"`
	MemberNickName string `json:"memberNickName"`
	Role           int8   `json:"role"`     //角色: 0-普通成员; 1-管理员; 2-群主
	IsMute         bool   `json:"isMute"`   //是否禁言
	MuteTime       int64  `json:"muteTime"` //禁言时间
}

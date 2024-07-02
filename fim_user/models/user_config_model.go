package models

import "gorm.io/gorm"

type UserConfigModel struct {
	gorm.Model
	UserID             uint            `json:"userID"`
	Online             bool            `json:"online"`             //是否在线
	RecallMsg          *string         `json:"recallMsg"`          //撤回消息的提示内容
	FriendOnlineNotify bool            `json:"friendOnlineNotify"` //好友上线通知
	Mute               bool            `json:"mute"`               //是否静音
	SecureLink         bool            `json:"secureLink"`         //是否使用安全链接
	SavePwd            bool            `json:"savePwd"`            //是否保存密码
	SearchUser         int8            `json:"searchUser"`         //是否允许搜索用户:0-no; 1-IDsearch; 2-Nicknamesearch
	FriendVerify       int8            `json:"friendVerify"`       //好友验证方式:0-no; 1-need verifyMSG; 2-need answer; 3-need answer correct; 4-allow everyone
	VerifyQuestion     *VerifyQuestion `json:"verifyQuestion"`     //验证问题: Only needed when friendVerify=2 or 3
}
type VerifyQuestion struct {
	Q1 *string `json:"q1"`
	A1 *string `json:"a1"`
	Q2 *string `json:"q2"`
	A2 *string `json:"a2"`
	Q3 *string `json:"q3"`
	A3 *string `json:"a3"`
}

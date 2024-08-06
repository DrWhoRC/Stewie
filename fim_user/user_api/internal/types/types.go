// Code generated by goctl. DO NOT EDIT.
package types

type FriendInfoRequest struct {
	UserId   uint `header:"UserId"`
	FriendId uint `json:"friendId"` //好友id
}

type FriendInfoResponse struct {
	Id       uint   `json:"id"`
	Nickname string `json:"nickname"`
	Role     int8   `json:"role"`
	Abstract string `json:"abstract"`
	Avatar   string `json:"avatar"`
	Notice   string `json:"notice"`
}

type FriendListRequest struct {
	UserId uint `header:"UserId"`
	Role   int8 `header:"Role"`
	Page   int  `form:"page"`
	Limit  int  `form:"limit"`
}

type FriendListResponse struct {
	List  []FriendInfoResponse `json:"list"`
	Count int                  `json:"count"`
}

type FriendNoticeUpdateRequest struct {
	UserId   uint   `header:"UserId"`
	FriendId uint   `json:"friendId"` //好友id
	Notice   string `json:"notice"`   //好友备注
}

type FriendNoticeUpdateResponse struct {
	Data string `json:"data"`
}

type SearchInfo struct {
	Id       uint   `json:"id"`
	Nickname string `json:"nickname"`
	Abstract string `json:"abstract"`
	Avatar   string `json:"avatar"`
	IsFriend bool   `json:"isFriend"` //is friend or not
}

type SearchRequest struct {
	UserId uint   `header:"UserId"`
	Key    string `form:"key"`    //id and nickname of users
	Online bool   `form:"online"` //online users
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
}

type SearchResponse struct {
	List  []SearchInfo `json:"list"`
	Count int          `json:"count"`
}

type UserConfResponse struct {
	UserId             uint           `header:"UserId"`
	Online             bool           `json:"online,optional" user_conf:"online"`                         //是否在线
	RecallMsg          string         `gorm:"size:32" json:"recallMsg,optional" user_conf:"recallMsg"`    //撤回消息的提示内容
	FriendOnlineNotify bool           `json:"friendOnlineNotify,optional" user_conf:"friendOnlineNotify"` //好友上线通知
	Mute               bool           `json:"mute,optional" user_conf:"mute"`                             //是否静音
	SecureLink         bool           `json:"secureLink,optional" user_conf:"secureLink"`                 //是否使用安全链接
	SavePwd            bool           `json:"savePwd,optional" user_conf:"savePwd"`                       //是否保存密码
	SearchUser         int8           `json:"searchUser,optional" user_conf:"searchUser"`                 //是否允许搜索用户:0-no; 1-IDsearch; 2-Nicknamesearch
	Verification       int8           `json:"verification,optional" user_conf:"verification"`             //好友验证方式:0-no; 1-need verifyMSG; 2-need answer; 3-need answer correct; 4-allow everyone
	VerifyQuestion     VerifyQuestion `json:"verifyQuestion,optional" user_conf:"verifyQuestion"`         //验证问题: Only needed when friendVerify=2 or 3
}

type UserConfUpdateRequest struct {
	UserId             uint            `header:"UserId"`
	Online             *bool           `json:"online,optional" user_conf:"online"`                         //是否在线
	RecallMsg          *string         `gorm:"size:32" json:"recallMsg,optional" user_conf:"recallMsg"`    //撤回消息的提示内容
	FriendOnlineNotify *bool           `json:"friendOnlineNotify,optional" user_conf:"friendOnlineNotify"` //好友上线通知
	Mute               *bool           `json:"mute,optional" user_conf:"mute"`                             //是否静音
	SecureLink         *bool           `json:"secureLink,optional" user_conf:"secureLink"`                 //是否使用安全链接
	SavePwd            *bool           `json:"savePwd,optional" user_conf:"savePwd"`                       //是否保存密码
	SearchUser         *int8           `json:"searchUser,optional" user_conf:"searchUser"`                 //是否允许搜索用户:0-no; 1-IDsearch; 2-Nicknamesearch
	Verification       *int8           `json:"verification,optional" user_conf:"verification"`             //好友验证方式:0-no; 1-need verifyMSG; 2-need answer; 3-need answer correct; 4-allow everyone
	VerifyQuestion     *VerifyQuestion `json:"verifyQuestion,optional" user_conf:"verifyQuestion"`         //验证问题: Only needed when friendVerify=2 or 3
}

type UserConfUpdateResponse struct {
	Data string `json:"data"`
}

type UserInfoRequest struct {
	UserId uint `header:"UserId"`
}

type UserInfoResponse struct {
	Id       uint   `json:"id"`
	Nickname string `json:"nickname"`
	Role     int8   `json:"role"`
	Abstract string `json:"abstract"`
	Avatar   string `json:"avatar"`
}

type UserInfoUpdateRequest struct {
	UserId   uint    `header:"UserId"`
	Nickname *string `json:"nickname,optional"`
	Role     *int8   `json:"role,optional"`
	Abstract *string `json:"abstract,optional"`
	Avatar   *string `json:"avatar,optional"`
}

type UserInfoUpdateResponse struct {
	Data string `json:"data"`
}

type UserValidRequest struct {
	UserId   uint   `header:"UserId"`
	FriendId uint   `json:"friendId"`
	ValidMsg string `json:"validMsg"`
}

type UserValidResponse struct {
	Verification   int8           `json:"verification"`
	VerifyQuestion VerifyQuestion `json:"verifyQuestion"`
	Data           string         `json:"data"`
}

type VerifyQuestion struct {
	Q1 *string `json:"q1"`
	A1 *string `json:"a1"`
	Q2 *string `json:"q2"`
	A2 *string `json:"a2"`
	Q3 *string `json:"q3"`
	A3 *string `json:"a3"`
}

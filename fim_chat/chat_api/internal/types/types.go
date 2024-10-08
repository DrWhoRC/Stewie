// Code generated by goctl. DO NOT EDIT.
package types

type ChatDeleteRequest struct {
	UserId uint `header:"UserId"`
	ChatId uint `form:"chatId"`
}

type ChatDeleteResponse struct {
	Data string `json:"data"`
}

type ChatHistoryRequest struct {
	UserId uint `header:"UserId"`
	Page   uint `form:"page"`
	Limit  uint `form:"limit"`
}

type ChatHistoryResponse struct {
	MsgId     uint   `json:"msgId"`
	UserId    uint   `json:"userId"`
	Avatar    string `json:"avatar"`
	Nickname  string `json:"nickname"`
	CreatedAt string `json:"createdAt"`
}

type ChatPinRequest struct {
	UserId   uint `header:"UserId"`
	FriendId uint `form:"friendId"`
}

type ChatPinResponse struct {
	Data string `json:"data"`
}

type ChatRequest struct {
	UserId uint `header:"UserId"`
}

type ChatResponse struct {
}

type ChatSessionDisplay struct {
	UserId     uint   `json:"userId"`
	Avatar     string `json:"avatar"`
	Nickname   string `json:"nickname"`
	CreatedAt  string `json:"created_at"`
	MsgPreview string `json:"msgPreview"`
}

type ChatSessionDisplayRequest struct {
	UserId uint `header:"UserId"`
	Page   uint `form:"page"`
	Limit  uint `form:"limit"`
	Key    int  `form:"key"`
}

type ChatSessionDisplayResponse struct {
	List  []ChatSessionDisplay `json:"list"`
	Count int                  `json:"count"`
}

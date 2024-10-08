package models

import (
	"fim/common/models/ctype"

	"gorm.io/gorm"
)

type GroupModel struct {
	gorm.Model
	Title string `gorm:"size:64" json:"title"`
	// TotalAmount int `json:"totalAmount"`
	// OnlineAmount int `json:"onlineAmount"`
	Abstract       string                `gorm:"size:128" json:"abstract"`
	Avatar         string                `gorm:"size:256" json:"avatar"`
	CreatorID      uint                  `json:"creatorID"`
	IsSearchable   bool                  `json:"isSearchable"`
	Verification   int8                  `json:"verification"` //验证方式: 0-no; 1-need verifyMSG; 2-need answer; 3-need answer correct; 4-allow everyone
	VerifyQuestion *ctype.VerifyQuestion `json:"verifyQuestion"`
	IsInvite       bool                  `json:"isInvite"`
	IsTemporary    bool                  `json:"isTemporary"`
	IsMute         bool                  `json:"isMute"` //全员禁言
	Size           int8                  `json:"size"`   //群规模 20,50,100,200,500,1000,2000,no limit
}

package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Pwd      string `gorm:"size:64" json:"pwd"`
	NickName string `gorm:"size:32" json:"nickname"`
	Abstract string `gorm:"size:128" json:"abstract"`
	Avatar   string `gorm:"size:256" json:"avatar"`
	IP       string `gorm:"size:32" json:"ip"`
	Address  string `gorm:"size:64" json:"address"`
}

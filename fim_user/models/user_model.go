package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Pwd      string `json:"pwd"`
	NickName string `json:"nickname"`
	Abstract string `json:"abstract"`
	Avatar   string `json:"avatar"`
	IP       string `json:"ip"`
	Address  string `json:"address"`
}

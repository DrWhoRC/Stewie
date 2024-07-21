package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Pwd             string           `gorm:"size:64" json:"-"`
	UserConfigModel *UserConfigModel `gorm:"foreignKey:UserID" json:"UserConfigModel"`
	Salt            string           `gorm:"size:32" json:"salt"`
	PwdWithSalt     string           `sql:"-" json:"pwdWithSalt"`
	NickName        string           `gorm:"size:32" json:"nickname"`
	Abstract        string           `gorm:"size:128" json:"abstract"`
	Avatar          string           `gorm:"size:256" json:"avatar"`
	IP              string           `gorm:"size:32" json:"ip"`
	Address         string           `gorm:"size:64" json:"address"`
	Role            int8             `gorm:"default:1" json:"role"` //1:user, 2:admin
}

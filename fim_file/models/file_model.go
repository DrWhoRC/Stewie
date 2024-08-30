package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileModel struct {
	gorm.Model
	Uid      uuid.UUID `json:"uid"`
	UserId   uint      `json:"userId"`
	FileName string    `json:"fileName"`
	FileSize int64     `json:"fileSize"`
	FilePath string    `json:"filePath"`
	Md5      string    `json:"md5"`
}

func (file *FileModel) Webpath() string {
	return "api/file/" + file.Uid.String()
}

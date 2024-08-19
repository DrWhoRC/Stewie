package models

import "gorm.io/gorm"

type FileModel struct {
	gorm.Model
	UserId   uint   `json:"userId"`
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
	FilePath string `json:"filePath"`
	WebPath  string `json:"webPath"`
	Md5      string `json:"md5"`
}

func (file *FileModel) Webpath() string {
	return "api/file/" + file.FilePath
}

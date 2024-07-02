package models

import (
	"fim/common/models/ctype"

	"gorm.io/gorm"
)

type ChatModel struct {
	gorm.Model
	SenderID   uint      `json:"senderID"`
	ReceiverID uint      `json:"receiverID"`
	MsgType    int8      `json:"msgType"`    //消息类型: 0-文本; 1-图片; 2-文件; 3-音频; 4-视频; 6-语音通话; 7-视频通话; 8-撤回消息; 9-引用消息;
	MsgPreview string    `json:"msgPreview"` //消息预览
	Msg        ctype.Msg `json:"msg"`        //消息内容
}

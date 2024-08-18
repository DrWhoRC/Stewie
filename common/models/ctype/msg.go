package ctype

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type MsgType int8

const (
	ImageMsgType MsgType = iota + 1
	VideoMsgType
	FileMsgType
	VoiceMsgType
	VoiceCallMsgType
	VideoCallMsgType
	WithdrawMsgType
	QuoteMsgType
	AtMsgType
)

type Msg struct {
	Type         int8          `json:"type"`    //消息类型: 0-文本; 1-图片; 2-文件; 3-音频; 4-视频; 6-语音通话; 7-视频通话; 8-撤回消息; 9-引用消息; 10-@
	Content      *string       `json:"content"` //only used when MsgType=0
	ImageMsg     *ImageMsg     `json:"imageMsg"`
	VideoMsg     *VideoMsg     `json:"videoMsg"`
	FileMsg      *FileMsg      `json:"fileMsg"`
	VoiceMsg     *VoiceMsg     `json:"voiceMsg"`
	VoiceCallMsg *VoiceCallMsg `json:"voiceCallMsg"`
	VideoCallMsg *VideoCallMsg `json:"videoCallMsg"`
	WithdrawMsg  *WithdrawMsg  `json:"withdrawMsg"`
	QuoteMsg     *QuoteMsg     `json:"quoteMsg"`
	AtMsg        *AtMsg        `json:"atMsg"`
}

func (c *Msg) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), c)
}
func (c Msg) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

type ImageMsg struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}
type VideoMsg struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Time  int    `json:"time"` //视频时长
}
type FileMsg struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Size  int64  `json:"size"` //文件大小
	Type  string `json:"type"` //文件类型
}
type VoiceMsg struct {
	URL  string `json:"url"`
	Time int    `json:"time"` //语音时长
}
type VoiceCallMsg struct {
	StartTime   time.Time `json:"startTime"`   //通话开始时间
	EndTime     time.Time `json:"endTime"`     //通话结束时间
	CallEnderID uint      `json:"callEnderID"` //通话结束者
}
type VideoCallMsg struct {
	StartTime   time.Time `json:"startTime"`   //通话开始时间
	EndTime     time.Time `json:"endTime"`     //通话结束时间
	CallEnderID uint      `json:"callEnderID"` //通话结束者
}
type WithdrawMsg struct {
	Content string `json:"content"` //撤回提示词
	Origin  *Msg   `json:"origin"`  //被撤回的消息
}
type QuoteMsg struct {
	MsgID uint `json:"msgID"` //引用的消息ID
	Msg   *Msg `json:"msg"`   //引用的消息
}
type AtMsg struct {
	AtUserID uint   `json:"atUserID"` //仅群聊中使用,被@的用户ID
	Content  string `json:"content"`
}

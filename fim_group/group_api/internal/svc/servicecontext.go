package svc

import (
	"fim/core"
	"fim/fim_chat/chat_rpc/chat"
	"fim/fim_group/group_api/internal/config"
	"fim/fim_user/user_rpc/users"

	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc users.Users
	ChatRpc chat.Chat
	DB      *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbconn := core.InitMysql()
	return &ServiceContext{
		Config:  c,
		UserRpc: users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
		ChatRpc: chat.NewChat(zrpc.MustNewClient(c.ChatRpc)),
		DB:      dbconn,
	}
}

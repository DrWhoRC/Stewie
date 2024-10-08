package svc

import (
	"fim/core"
	"fim/fim_chat/chat_api/internal/config"
	"fim/fim_chat/chat_rpc/chat"
	"fim/fim_user/user_rpc/users"

	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc users.Users
	ChatRpc chat.Chat
	Redis   *redis.Client
	DB      *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbconn := core.InitMysql()
	redisClient := core.InitRedis(c.Redis.Addr, c.Redis.Pwd, c.Redis.DB)
	return &ServiceContext{
		Config:  c,
		UserRpc: users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
		ChatRpc: chat.NewChat(zrpc.MustNewClient(c.ChatRpc)),
		Redis:   redisClient,
		DB:      dbconn,
	}
}

package svc

import (
	"fim/core"
	"fim/fim_auth/auth_api/internal/config"
	"fim/fim_user/user_rpc/users"

	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Redis   *redis.Client
	UserRpc users.Users
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbconn := core.InitMysql()
	redisClient := core.InitRedis(c.Redis.Addr, c.Redis.Pwd, c.Redis.DB)
	return &ServiceContext{
		Config:  c,
		DB:      dbconn,
		Redis:   redisClient,
		UserRpc: users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
	}
}

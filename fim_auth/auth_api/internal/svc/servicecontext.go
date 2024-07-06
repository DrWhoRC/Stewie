package svc

import (
	"fim/core"
	"fim/fim_auth/auth_api/internal/config"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbconn := core.InitMysql()
	redisClient := core.InitRedis(c.Redis.Addr, c.Redis.Pwd, c.Redis.DB)
	return &ServiceContext{
		Config: c,
		DB:     dbconn,
		Redis:  redisClient,
	}
}

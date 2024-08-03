package svc

import (
	"fim/core"
	"fim/fim_user/user_api/internal/config"
	"fim/fim_user/user_rpc/users"

	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc users.Users
	DB      *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbconn := core.InitMysql()
	return &ServiceContext{
		Config:  c,
		UserRpc: users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
		DB:      dbconn,
	}
}

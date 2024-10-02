package svc

import (
	"fim/core"
	"fim/fim_group/group_rpc/internal/config"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbconn := core.InitMysql()
	return &ServiceContext{
		Config: c,
		DB:     dbconn,
	}
}

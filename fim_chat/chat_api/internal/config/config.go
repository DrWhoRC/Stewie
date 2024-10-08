package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}

	Redis struct {
		Addr string
		Pwd  string
		DB   int
	}
	UserRpc zrpc.RpcClientConf
	ChatRpc zrpc.RpcClientConf
	Etcd    string
	Auth    struct {
		AccessSecret string
		AccessExpire int
	}
}

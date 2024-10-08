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
	Auth struct {
		AccessSecret string
		AccessExpire int
	}
	Redis struct {
		Addr string
		Pwd  string
		DB   int
	}
	ThirdPartyLoginList []struct {
		Name string
		Icon string
		Url  string
	}
	QQ struct {
		AppID    string
		AppKey   string
		Redirect string
	}
	UserRpc zrpc.RpcClientConf
	Etcd    string
}

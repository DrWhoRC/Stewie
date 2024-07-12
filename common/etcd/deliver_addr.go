package etcd

import (
	"context"
	"fim/core"
	"fim/utils/ips"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

func DeliverAddr(etcdAddr string, serviceName string, addr string) {
	list := strings.Split(addr, ":")
	if len(list) != 2 {
		logx.Error("addr format error")
		return
	}
	ip := list[0]
	if ip == "0.0.0.0" {
		ip = ips.GetIP()
		addr = strings.ReplaceAll(addr, "0.0.0.0", ip)
	}

	client, err := core.InitEtcd(etcdAddr)
	if err != nil {
		logx.Error("etcd init error", err)
		return

	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.Put(ctx, serviceName, addr)

	if err != nil {
		logx.Error("地址上送失败", err)
	}
	logx.Infof("地址上送成功:%s,%s", serviceName, addr)
	fmt.Printf("地址上送成功:%s,%s", serviceName, addr)
}

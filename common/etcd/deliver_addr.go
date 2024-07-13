package etcd

import (
	"context"
	"fim/core"
	"fim/utils/ips"
	"fmt"
	"strings"

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

	_, err = client.Put(context.Background(), serviceName, addr)

	if err != nil {
		logx.Error("地址上送失败", err)
	}
	logx.Infof("地址上送成功:%s,%s", serviceName, addr)
	fmt.Printf("地址上送成功:%s,%s", serviceName, addr)
}

func GetServiceAddr(EtcdAddr string, ServiceName string) string {
	client, err := core.InitEtcd(EtcdAddr)
	if err != nil {
		logx.Error("etcd init error", err)
		return "etcd init error"
	}
	addr, err := client.Get(context.Background(), ServiceName)
	if err == nil && len(addr.Kvs) > 0 {
		return string(addr.Kvs[0].Value)
	}
	return "get addr error"

}

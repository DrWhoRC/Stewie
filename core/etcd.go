package core

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func InitEtcd(addr string) (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {

		return nil, err
	}
	return cli, nil
}

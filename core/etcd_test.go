package core

import (
	"context"
	"fmt"
	"testing"
)

func TestEtcd(t *testing.T) {
	client, _ := InitEtcd("localhost:2379")
	res, err := client.Put(context.Background(), "auth_api", "127.0.0.1:8888")
	fmt.Println(res, err)
	getResponse, err := client.Get(context.Background(), "auth_api")
	fmt.Println(getResponse, err)
}

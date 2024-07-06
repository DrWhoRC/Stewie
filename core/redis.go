package core

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func InitRedis(addr, pwd string, db int) (client *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd, // no password set
		DB:       db,  // use default DB
		PoolSize: 100,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		panic("failed to connect redis")
	}
	return rdb

}

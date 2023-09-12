package redis

import (
	redsync "github.com/go-redsync/redsync/v4"
	redis2 "github.com/go-redsync/redsync/v4/redis"
	goredis "github.com/go-redsync/redsync/v4/redis/goredis/v9"
	redis "github.com/redis/go-redis/v9"
)

func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func NewPool(client *redis.Client) redis2.Pool {
	return goredis.NewPool(client)
}

func NewRedsync(pool redis2.Pool) *redsync.Redsync {
	return redsync.New(pool)
}

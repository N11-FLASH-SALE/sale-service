package redis

import (
	"sale/config"

	"github.com/redis/go-redis/v9"
)

func ConnectDB() *redis.Client {
	conf := config.Load()
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.RDB_ADDRESS,
		Password: conf.Redis.RDB_PASSWORD,
		DB:       0,
	})

	return rdb
}

package config

import "github.com/redis/go-redis/v9"

func BuildCache() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return rdb
}

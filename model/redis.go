package model

import (
	"github.com/go-redis/redis/v7"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := RedisClient.Ping().Err()
	if err != nil {
		panic(err)
	}
}

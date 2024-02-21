package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
)

var redisClient *redis.Client
var once sync.Once
var ctx = context.Background()

func GetClient() *redis.Client {
	once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     "redis-18155.c281.us-east-1-2.ec2.cloud.redislabs.com:18155",
			Password: "chMtp0E9VmLtdsphdKQpgX81DMCHZwwM", // no password set
			DB:       0,                                  // use default DB
		})
	})

	return redisClient
}

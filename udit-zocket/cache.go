package cache

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379", // Redis address
})

func SetCache(key string, value string) {
	ctx := context.Background()
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Printf("Error setting cache: %v", err)
	}
}

func GetCache(key string) (string, error) {
	ctx := context.Background()
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		log.Printf("Error reading from cache: %v", err)
	}
	return val, err
}

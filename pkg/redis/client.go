package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var ctx = context.Background()

func Connect(host, port, password string, db int) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})

	err := Client.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	log.Println("Redis connected successfully")
	return nil
}

func Set(key string, value interface{}, expiration time.Duration) error {
	return Client.Set(ctx, key, value, expiration).Err()
}

func Get(key string) (string, error) {
	return Client.Get(ctx, key).Result()
}

func Delete(key string) error {
	return Client.Del(ctx, key).Err()
}

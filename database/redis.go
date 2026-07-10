package database

import (
	"context"
	"fmt"
	"os"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func ConnectRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		),
		Password: "", // No password set
		DB: 0, // Use default DB
	})
	_,err := rdb.Ping(Ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis successfully")
	return rdb, nil
}
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func Connect(redisURL string) *redis.Client {
	var opts *redis.Options
	var err error
	if redisURL != "" {
		opts, err = redis.ParseURL(redisURL)
		if err != nil {
			panic(err)
		}
	} else {
		opts = &redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}
	}
	return redis.NewClient(opts)
}

func NewServer() {
	rdb := Connect("")

	err := rdb.Set(ctx, "key", "value", 15*time.Second).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}

func main() {
	NewServer()
}

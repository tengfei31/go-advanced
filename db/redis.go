package db

import (
	"log"
	"time"
	"github.com/go-redis/redis"
)

var Redis *redis.Client

// var ctx =  context.Background()

func init() {
	if err := initRedis(); err != nil {
		log.Fatal(err)
	}
}

func initRedis() error {
	Redis = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	_, err := Redis.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

package rds

import "github.com/redis/go-redis/v9"

var Rds *redis.Client

func InitRedis() (err error) {
	Rds = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
		PoolSize: 10,
	})
	return err
}

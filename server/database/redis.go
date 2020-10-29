package database

import (
	"github.com/go-redis/redis"
)

var (
	RDB *redis.Client
)

func InitRedis() error {
	RDB := redis.NewClient(&redis.Options{
		Addr:     RedisConnectAddr,
		Password: RedisPassword,
		DB:       RedisDatabase,
	})

	_, err := RDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

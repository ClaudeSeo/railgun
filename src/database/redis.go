package database

import (
	"log"

	"github.com/claudeseo/railgun/src/config"
	"github.com/go-redis/redis"
)

var conn *redis.Client

func Init() {
	cfg := config.GetConfig()
	conn = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.ReidsPassword,
		DB:       cfg.RedisDatabase,
	})

	if _, err := conn.Ping().Result(); err != nil {
		log.Fatal("Can't Redis Ping ", cfg.RedisAddr)
	}
}

func GetRedis() *redis.Client {
	return conn
}

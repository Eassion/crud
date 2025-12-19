package cache

import (
	"context"
	"fmt"
	"log"

	"crud/config"

	"github.com/redis/go-redis/v9"
)

var (
	RDB *redis.Client
	Ctx = context.Background()
)

func InitRedis() {
	cfg := config.Conf.Redis

	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := RDB.Ping(Ctx).Err(); err != nil {
		log.Fatal("Redis 连接失败:", err)
	}

	log.Println("Redis 连接成功")
}

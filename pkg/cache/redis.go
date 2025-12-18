package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	RDB *redis.Client
	Ctx = context.Background()
)

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // 有就填
		DB:       0,
	})

	if err := RDB.Ping(Ctx).Err(); err != nil {
		log.Fatal("Redis 连接失败:", err)
	}

	log.Println("Redis 连接成功")
}

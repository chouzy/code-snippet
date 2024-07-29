package redis

import (
	"code-snippet/pkg/setting"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// InitRedis redis初始化
func InitRedis(r *setting.RedisSettingS) *redis.Client {
	// 普通连接方式
	opt := redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
		PoolSize: r.PoolSize, // 数据库连接池大小
	}
	// 解析URL：opt = redis.ParseURL("redis://<user>:<pass>@localhost:6379/<mysql>")
	client := redis.NewClient(&opt)
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis init failed, err: %v \n", err))
	}
	fmt.Printf("redis connect ping response: %v \n", pong)
	return client
}

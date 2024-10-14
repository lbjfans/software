package redis

import (
	"backend/settings"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host, cfg.Port,
		),
		//Addr:     "localhost:6379",
		Password: viper.GetString("redis.password"), // 密码
		DB:       viper.GetInt("redis.db"),          // 数据库
		PoolSize: viper.GetInt("redis.poll_size"),   // 连接池大小
	})
	_, err = rdb.Ping(context.Background()).Result()
	return
}

func Close() {
	rdb.Close()
}

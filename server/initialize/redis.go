package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"server/global"
)

func Redis() {
	redisConfig := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Log.Error("redis连接失败,err:", zap.Error(err))
	} else {
		global.Log.Info("redis连接成功,pong:", zap.String("pong", pong))
		global.Redis = client
	}
}

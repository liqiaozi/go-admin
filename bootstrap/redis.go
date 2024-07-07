package bootstrap

import (
	"context"
	"github.com/redis/go-redis/v9"
	"lixuefei.com/go-admin/global"
	"lixuefei.com/go-admin/global/logger"
)

func initRedis() {
	var client redis.UniversalClient

	redisConfig := global.App.Server.Redis
	if redisConfig.UseCluster {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    redisConfig.ClusterAddress,
			Password: redisConfig.Password,
		})
	} else {
		// 使用单例模式
		client = redis.NewClient(&redis.Options{
			Addr:     redisConfig.Address,
			Password: redisConfig.Password,
			DB:       redisConfig.DB,
		})
	}

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Log.Errorf("redis connect ping failed, err: %v", err)
		panic(err)
	} else {
		logger.Log.Infof("redis connect ping response:: %v", pong)
		global.App.RedisClient = &client
	}
}

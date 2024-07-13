package bootstrap

import (
	"context"
	"github.com/redis/go-redis/v9"
	"lixuefei.com/go-admin/common/component/logger"
	"lixuefei.com/go-admin/global"
)

func initRedis() {
	logger.Log.Infof("[bootstrap] init redis begin...")
	var client redis.UniversalClient
	redisConfig := global.App.Server.Redis
	if redisConfig.UseCluster {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    redisConfig.ClusterAddress,
			Password: redisConfig.Password,
		})
		logger.Log.Infof("[bootstrap] redis mode: cluster")
	} else {
		// 使用单例模式
		client = redis.NewClient(&redis.Options{
			Addr:     redisConfig.Address,
			Password: redisConfig.Password,
			DB:       redisConfig.DB,
		})
		logger.Log.Infof("[bootstrap] redis mode: standalone")
	}

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Log.Errorf("[bootstrap] redis connect ping failed, err: %v", err)
		panic(err)
	} else {
		logger.Log.Infof("[bootstrap] redis connect ping response: %v", pong)
		global.App.RedisClient = client
		logger.Log.Infof("[bootstrap] init redis end...")
	}
}

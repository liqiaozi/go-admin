package bootstrap

import (
	"lixuefei.com/go-admin/common/component/logger"
	"lixuefei.com/go-admin/global"
)

func DestroyProcesses() {
	if global.App.DB != nil {
		db, _ := global.App.DB.DB()
		db.Close()
		logger.Log.Infof("[destroy] close database...")
	}

	if global.App.RedisClient != nil {
		redisClient := global.App.RedisClient
		redisClient.Close()
		logger.Log.Infof("[destroy] close redis...")
	}
}

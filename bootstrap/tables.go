package bootstrap

import (
	"lixuefei.com/go-admin/global"
	"lixuefei.com/go-admin/global/logger"
	"lixuefei.com/go-admin/internal/models"
	"os"
)

func registerTable() {
	db := global.App.DB
	err := db.AutoMigrate(
		models.SysUserEntity{},
	)
	if err != nil {
		logger.Log.Errorf("register table error: %s", err.Error())
		os.Exit(0)
	}
	logger.Log.Infof("register table success...")
}

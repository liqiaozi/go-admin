package bootstrap

import (
	"lixuefei.com/go-admin/app/admin/model"
	"lixuefei.com/go-admin/common/component/logger"
	"lixuefei.com/go-admin/global"
	"os"
)

func registerTable() {
	logger.Log.Infof("[bootstrap] register table begin...")
	db := global.App.DB
	err := db.AutoMigrate(
		model.SysRole{},
		model.SysUser{},
		model.SysUserRole{},
		model.SysMenu{},
		model.SysRoleMenu{},
	)
	if err != nil {
		logger.Log.Errorf("register table error: %s", err.Error())
		os.Exit(0)
	}
	logger.Log.Infof("[bootstrap] register table end...")
}

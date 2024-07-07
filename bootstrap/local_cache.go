package bootstrap

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"lixuefei.com/go-admin/global"
)

func initLocalCache() {
	global.App.BlackCache = local_cache.NewCache()
}

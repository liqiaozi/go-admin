package router

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/common/component/logger"
	"lixuefei.com/go-admin/common/middleware"
	"lixuefei.com/go-admin/router/admin"
)

// InitRouter 路由初始化
func InitRouter() *gin.Engine {
	logger.Log.Infof("[router] init router begin...")
	r := gin.New()
	r.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		r.Use(gin.Logger())
	}

	r.Use(middleware.CustomExceptionHandler())
	// 系统管理应用路由初始化
	admin.InitRouter(r)

	logger.Log.Infof("[router] init router end...")
	return r
}

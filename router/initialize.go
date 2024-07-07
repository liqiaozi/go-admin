package router

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/global/logger"
	"net/http"
)

// 路由初始化
func InitializeRouter() *gin.Engine {
	logger.Log.Infof("[router] init router begin...")
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	PublicGroup := Router.Group("/go-admin/open")
	PrivateGroup := Router.Group("/go-admin/api")
	//PrivateGroup.Use(middleware.JWTAuth())

	// OpenAPI
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	systemRouter := RouterGroupApp.System
	{
		systemRouter.InitSysUserRouter(PrivateGroup) // 注册用户路由
	}
	logger.Log.Infof("[router] init router end...")
	return Router
}

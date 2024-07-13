package router

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/common/component/logger"
	"lixuefei.com/go-admin/router/admin"
	"net/http"
)

type RouterGroup struct {
	Admin admin.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

// 路由初始化
func InitRouter() *gin.Engine {
	logger.Log.Infof("[router] init router begin...")
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	PublicGroup := Router.Group("/go-admin/api")
	//PrivateGroup.Use(middleware.JWTAuth())

	// 不做鉴权的路由
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	adminRouter := RouterGroupApp.Admin
	{
		adminRouter.InitRoleRouter(PublicGroup)
		adminRouter.InitUserRouter(PublicGroup)
		adminRouter.InitCaptchaRouter(PublicGroup)
	}
	logger.Log.Infof("[router] init router end...")
	return Router
}

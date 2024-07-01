package main

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/bootstrap"
	"lixuefei.com/go-admin/global"
	"lixuefei.com/go-admin/global/logger"
	"net/http"
)

func main() {
	bootstrap.Init()
	logger.Log.Info("aaa")
	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 启动服务器
	port := global.App.Application.AppConfig.Port
	logger.Log.Errorf("服务成功启动，运行在端口：%v", port)
	r.Run(":" + port)

}

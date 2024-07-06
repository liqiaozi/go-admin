package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/bootstrap"
	"lixuefei.com/go-admin/global"
	"lixuefei.com/go-admin/global/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 初始化操作
	bootstrap.Init()

	// 创建http服务器
	r := gin.Default()
	srv := &http.Server{
		Addr:    ":" + global.App.Server.AppConfig.Port,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Infof("listen: %s\n", err)
		}
	}()

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Infof("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Infof("Server Shutdown:", err)
	}
	logger.Log.Infof("Server exiting")
}

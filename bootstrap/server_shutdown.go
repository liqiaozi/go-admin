package bootstrap

import (
	"context"
	"lixuefei.com/go-admin/global/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ServerShutdown(srv *http.Server) {
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Infof("[server] server shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Errorf("[server] server shutdown error: %s", err)
	}
	logger.Log.Infof("[server] server exiting")
}

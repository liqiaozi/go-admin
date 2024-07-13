package bootstrap

import (
	"errors"
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/common/component/logger"
	"lixuefei.com/go-admin/global"
	"net/http"
	"time"
)

func InitServer(address string, router *gin.Engine) *http.Server {
	logger.Log.Infof("[server] init server begin...")
	srv := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logger.Log.Infof("[server] server run on port: %s\n", global.App.Server.ServiceInfo.Port)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Errorf("[server] server listen err: %s\n", err)
		}
	}()
	return srv
}

package system

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/internal/controller"
)

type SysBaseRouter struct{}

func (s *SysBaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("/base")
	captchaController := controller.ControllerGroupApp.CaptchaController
	{
		baseRouter.GET("/captcha", captchaController.GenCaptcha)
	}
}

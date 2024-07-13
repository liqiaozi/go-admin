package admin

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/app/admin/controller"
)

type CaptchaRouter struct {
}

func (s *CaptchaRouter) InitCaptchaRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("/v1/captcha/")
	{
		userRouter.GET("/get", controller.CaptchaController{}.GenCaptcha)
	}
}

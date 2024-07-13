package admin

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/app/admin/controller"
)

type SysUserRouter struct {
}

func (s *SysUserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("/v1/user/")
	{
		userRouter.POST("/queryUserList", controller.SysUserController{}.QueryUserList)
	}
}

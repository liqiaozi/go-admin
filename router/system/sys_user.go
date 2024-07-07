package system

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/internal/controller"
)

type SysUserRouter struct{}

func (s *SysUserRouter) InitSysUserRouter(Router *gin.RouterGroup) {
	sysUserRouter := Router.Group("/user")
	userController := controller.ControllerGroupApp.SysUserController
	{
		sysUserRouter.POST("/queryUserList", userController.QueryUserList)
	}
}

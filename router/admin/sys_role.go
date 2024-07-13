package admin

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/app/admin/controller"
)

type SysRoleRouter struct {
}

func (s *SysUserRouter) InitRoleRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("/v1/role/")
	{
		userRouter.POST("/queryRoleList", controller.SysRoleController{}.QueryRoleList)
	}
}

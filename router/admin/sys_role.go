package admin

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/app/admin/controller"
)

func init() {
	checkRouter = append(checkRouter, registerSysRoleRouter)
}

func registerSysRoleRouter(v1 *gin.RouterGroup) {
	roleController := controller.SysRoleController{}
	r := v1.Group("/role")
	{
		r.POST("/add", roleController.AddSysRole)
		r.GET("/detail", roleController.GetSysRole)
		r.POST("/update", roleController.UpdateSysRole)
		r.POST("/updateStatus", roleController.UpdateSysRoleStatus)
		r.POST("/pageList", roleController.QuerySysRoleList)
		r.POST("/delete", roleController.DeleteSysRole)
	}
}

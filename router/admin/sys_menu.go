package admin

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/app/admin/controller"
)

func init() {
	checkRouter = append(checkRouter, registerSysMenuRouter)
}

func registerSysMenuRouter(v1 *gin.RouterGroup) {
	menuController := controller.SysMenuController{}
	r := v1.Group("/menu")
	{
		r.POST("/add", menuController.AddSysMenu)
		r.GET("/detail", menuController.GetSysMenuById)
		r.GET("/tree", menuController.GetSysMenuTree)
		r.POST("/update", menuController.UpdateSysMenu)
		r.POST("/delete", menuController.DeleteSysMenu)
	}
}

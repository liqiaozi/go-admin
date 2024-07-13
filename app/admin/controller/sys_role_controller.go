package controller

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/common/response"
)

type SysRoleController struct {
}

func (s SysRoleController) QueryRoleList(c *gin.Context) {
	c.JSON(200, response.Ok())
}

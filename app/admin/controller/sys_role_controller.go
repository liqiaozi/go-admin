package controller

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/common/response"
)

type SysRoleController struct {
}

func (s SysRoleController) AddSysRole(c *gin.Context) {
	c.JSON(200, response.Ok())
}

func (s SysRoleController) GetSysRole(c *gin.Context) {
	c.JSON(200, response.Ok())
}

func (s SysRoleController) UpdateSysRole(c *gin.Context) {
	c.JSON(200, response.Ok())
}

func (s SysRoleController) GetSysRoleList(c *gin.Context) {
	c.JSON(200, response.Ok())
}

func (s SysRoleController) DeleteSysRole(c *gin.Context) {
	c.JSON(200, response.Ok())
}

func (s SysRoleController) UpdateSysRoleStatus(c *gin.Context) {
	c.JSON(200, response.Ok())
}

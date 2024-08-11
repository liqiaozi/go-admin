package controller

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/app/admin/model/dto"
	"lixuefei.com/go-admin/app/admin/service"
	"lixuefei.com/go-admin/common/errors"
	"lixuefei.com/go-admin/common/response"
	"strconv"
)

type SysRoleController struct {
}

// AddSysRole 新增角色
func (s SysRoleController) AddSysRole(c *gin.Context) {
	var sysRoleDTO dto.SysRoleDTO
	err := c.ShouldBindJSON(sysRoleDTO)
	if err != nil {
		errors.ThrowExceptionWithMsg(errors.ParamsError, "新增角色参数异常")
		return
	}
	roleId := service.SysRoleService{}.AddSysRole(&sysRoleDTO)
	c.JSON(200, response.OkWithData(roleId))
}

// GetSysRole 查询角色详情
func (s SysRoleController) GetSysRole(c *gin.Context) {
	roleIdStr := c.Query("roleId")
	if roleIdStr == "" {
		errors.ThrowExceptionWithMsg(errors.ParamsError, "角色ID不存在")
		return
	}
	roleId, _ := strconv.Atoi(roleIdStr)
	sysRole := service.SysRoleService{}.QuerySysRoleById(roleId)
	c.JSON(200, response.OkWithData(sysRole))
}

func (s SysRoleController) UpdateSysRole(c *gin.Context) {
	c.JSON(200, response.Ok())
}

func (s SysRoleController) QuerySysRoleList(c *gin.Context) {
	c.JSON(200, response.Ok())
}

func (s SysRoleController) DeleteSysRole(c *gin.Context) {
	c.JSON(200, response.Ok())
}

func (s SysRoleController) UpdateSysRoleStatus(c *gin.Context) {
	c.JSON(200, response.Ok())
}

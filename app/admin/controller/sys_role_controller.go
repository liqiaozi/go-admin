package controller

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/app/admin/model/dto"
	"lixuefei.com/go-admin/app/admin/service"
	"lixuefei.com/go-admin/common/errors"
	"lixuefei.com/go-admin/common/response"
	"lixuefei.com/go-admin/common/utils"
)

type SysRoleController struct {
}

// AddSysRole 新增角色
func (s SysRoleController) AddSysRole(c *gin.Context) {
	var addSysRole dto.SysRoleDTO
	err := c.ShouldBindJSON(&addSysRole)
	if err != nil {
		errors.ThrowExceptionWithMsg(errors.ParamsError, "新增角色参数不合法")
		return
	}
	roleId := service.SysRoleService{}.AddSysRole(&addSysRole)
	c.JSON(200, response.OkWithData(roleId))
}

// GetSysRole 查询角色
func (s SysRoleController) GetSysRole(c *gin.Context) {
	roleIdStr := c.Query("roleId")
	if roleIdStr == "" {
		errors.ThrowExceptionWithMsg(errors.ParamsError, "角色ID参数不存在")
		return
	}
	roleId, _ := utils.StringToInt(roleIdStr)
	sysRole := service.SysRoleService{}.QuerySysRoleById(roleId)
	c.JSON(200, response.OkWithData(sysRole))
}

// UpdateSysRole 更新角色
func (s SysRoleController) UpdateSysRole(c *gin.Context) {
	var updateSysRole dto.SysRoleDTO
	err := c.ShouldBindJSON(&updateSysRole)
	if err != nil {
		errors.ThrowExceptionWithMsg(errors.ParamsError, "更新角色参数不合法")
		return
	}
	service.SysRoleService{}.UpdateSysRole(&updateSysRole)
	c.JSON(200, response.Ok())
}

// QuerySysRoleList 查询角色列表
func (s SysRoleController) QuerySysRoleList(c *gin.Context) {
	var rolePageQuery dto.SysRolePageQueryDTO
	err := c.ShouldBindJSON(&rolePageQuery)
	if err != nil {
		errors.ThrowExceptionWithMsg(errors.ParamsError, "查询角色列表参数不合法")
		return
	}
	pageResult := service.SysRoleService{}.QuerySysRoleListByPage(&rolePageQuery)
	c.JSON(200, response.OkWithData(pageResult))
}

// DeleteSysRole 删除角色
func (s SysRoleController) DeleteSysRole(c *gin.Context) {
	roleIdStr := c.Query("roleId")
	if roleIdStr == "" {
		errors.ThrowExceptionWithMsg(errors.ParamsError, "角色ID参数不存在")
		return
	}
	roleId, _ := utils.StringToInt(roleIdStr)
	service.SysRoleService{}.DeleteSysRoleById(roleId)
	c.JSON(200, response.Ok())
}

// UpdateSysRoleStatus 更新角色状态
func (s SysRoleController) UpdateSysRoleStatus(c *gin.Context) {
	var updateRoleStatusReq dto.UpdateRoleStatusReqDTO
	err := c.ShouldBindJSON(&updateRoleStatusReq)
	if err != nil {
		errors.ThrowExceptionWithMsg(errors.ParamsError, "更新角色状态参数不合法")
		return
	}
	service.SysRoleService{}.UpdateSysRoleStatus(updateRoleStatusReq.RoleId, updateRoleStatusReq.Status)
	c.JSON(200, response.Ok())
}

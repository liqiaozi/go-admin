package controller

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/app/admin/model/dto"
	"lixuefei.com/go-admin/app/admin/service"
	"lixuefei.com/go-admin/common/errors"
	"lixuefei.com/go-admin/common/request"
	"lixuefei.com/go-admin/common/response"
)

type SysUserController struct {
}

// AddSysUser 新增用户
func (s SysUserController) AddSysUser(c *gin.Context) {
	var register dto.SysUserRegister
	err := c.ShouldBindJSON(&register)
	if err != nil {
		c.JSON(200, response.FailByErrorWithMsg(errors.ParamsError, err.Error()))
		return
	}
	sysUserId := service.SysUserService{}.Create(register)
	c.JSON(200, response.OkWithData(sysUserId))
}

// GetSysUser 获取用户
func (s SysUserController) GetSysUser(c *gin.Context) {
	c.JSON(200, response.Ok())
}

// UpdateSysUser 更新用户
func (s SysUserController) UpdateSysUser(c *gin.Context) {
	c.JSON(200, response.Ok())
}

// QuerySysUserList 查询用户列表
func (s SysUserController) QuerySysUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		c.JSON(200, response.FailByError(errors.ParamsError))
	}
	result := service.SysUserService{}.QueryUserList(pageInfo)
	c.JSON(200, response.OkWithData(result))
}

// DeleteSysUser 删除用户
func (s SysUserController) DeleteSysUser(c *gin.Context) {
	c.JSON(200, response.Ok())
}

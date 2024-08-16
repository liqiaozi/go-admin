package controller

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/app/admin/model/dto"
	"lixuefei.com/go-admin/app/admin/service"
	"lixuefei.com/go-admin/common/errors"
	"lixuefei.com/go-admin/common/response"
	"strconv"
)

type SysUserController struct {
}

// AddSysUser 新增用户
func (s SysUserController) AddSysUser(c *gin.Context) {
	var register dto.SysUserRegisterDTO
	err := c.ShouldBindJSON(&register)
	if err != nil {
		errors.ThrowExceptionWithMsg(errors.ParamsError, "新建用户参数异常")
		return
	}
	sysUserId := service.SysUserService{}.AddSysUser(&register)
	c.JSON(200, response.OkWithData(sysUserId))
}

// GetSysUser 获取用户
func (s SysUserController) GetSysUser(c *gin.Context) {
	userIdStr := c.Query("userId")
	if userIdStr == "" {
		errors.ThrowExceptionWithMsg(errors.ParamsError, "用户ID参数不存在")
		return
	}
	userId, _ := strconv.Atoi(userIdStr)
	sysUser := service.SysUserService{}.QuerySysUserById(userId)
	c.JSON(200, response.OkWithData(sysUser))
}

// UpdateSysUser 更新用户
func (s SysUserController) UpdateSysUser(c *gin.Context) {
	var updateUser dto.SysUserUpdateDTO
	err := c.ShouldBindJSON(&updateUser)
	if err != nil {
		errors.ThrowExceptionWithMsg(errors.ParamsError, "更新用户参数异常")
		return
	}
	sysUserId := service.SysUserService{}.UpdateSysUser(&updateUser)
	c.JSON(200, response.OkWithData(sysUserId))
}

// DeleteSysUser 删除用户
func (s SysUserController) DeleteSysUser(c *gin.Context) {
	userIdStr := c.Query("userId")
	if userIdStr == "" {
		errors.ThrowExceptionWithMsg(errors.ParamsError, "用户ID参数不存在")
		return
	}
	userId, _ := strconv.Atoi(userIdStr)
	service.SysUserService{}.DeleteSysUserById(userId)
	c.JSON(200, response.Ok())
}

// UpdateUserStatus 更新用户状态
func (s SysUserController) UpdateUserStatus(c *gin.Context) {

	var userStatusUpdateDTO dto.SysUserStatusUpdateDTO
	err := c.ShouldBindJSON(&userStatusUpdateDTO)
	if err != nil {
		errors.ThrowExceptionWithMsg(errors.ParamsError, "更新用户状态异常")
		return
	}
	sysUserId := service.SysUserService{}.UpdateSysUserStatus(userStatusUpdateDTO.UserId, userStatusUpdateDTO.Status)
	c.JSON(200, response.OkWithData(sysUserId))
}

// QuerySysUserList 查询用户列表
func (s SysUserController) QuerySysUserList(c *gin.Context) {
	var userPageInfo dto.SysUserPageQueryDTO
	err := c.ShouldBindJSON(&userPageInfo)
	if err != nil {
		c.JSON(200, response.FailByError(errors.ParamsError))
	}
	result := service.SysUserService{}.QueryUserList(&userPageInfo)
	c.JSON(200, response.OkWithData(result))
}

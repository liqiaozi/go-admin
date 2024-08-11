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

// Create 新增用户
func (s SysUserController) Create(c *gin.Context) {
	var register dto.SysUserRegister
	err := c.ShouldBindJSON(&register)
	if err != nil {
		c.JSON(200, response.FailByErrorWithMsg(errors.ParamsError, err.Error()))
		return
	}
	sysUserId := service.SysUserService{}.Create(register)
	c.JSON(200, response.OkWithData(sysUserId))
}

func (s SysUserController) QueryUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		c.JSON(200, response.FailByError(errors.ParamsError))
	}
	result := service.SysUserService{}.QueryUserList(pageInfo)
	c.JSON(200, response.OkWithData(result))
}

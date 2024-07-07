package controller

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/common/errors"
	"lixuefei.com/go-admin/common/request"
	"lixuefei.com/go-admin/common/response"
)

type SysUserController struct {
}

func (u *SysUserController) QueryUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		c.JSON(200, response.FailByError(errors.ParamsError))
	}
	result := sysUserService.QueryUserList(pageInfo)
	c.JSON(200, response.OkWithData(result))
}

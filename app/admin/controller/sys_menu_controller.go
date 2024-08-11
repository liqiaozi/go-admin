package controller

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/app/admin/model/dto"
	"lixuefei.com/go-admin/app/admin/service"
	"lixuefei.com/go-admin/common/errors"
	"lixuefei.com/go-admin/common/response"
	"strconv"
)

type SysMenuController struct{}

// AddSysMenu 新增菜单
func (e SysMenuController) AddSysMenu(c *gin.Context) {
	sysMenuAddReq := dto.SysMenuAddReqDTO{}
	err := c.ShouldBindJSON(sysMenuAddReq)
	if err != nil {
		errors.ThrowException(errors.ParamsError)
		return
	}
	menuId := service.SysMenuService{}.AddSysMenu(&sysMenuAddReq)
	c.JSON(200, response.OkWithData(menuId))
}

// GetSysMenuById 查询菜单详情
func (e SysMenuController) GetSysMenuById(c *gin.Context) {
	menuIdStr := c.Query("menuId")
	if menuIdStr == "" {
		errors.ThrowException(errors.ParamsError)
		return
	}
	menuId, _ := strconv.Atoi(menuIdStr)
	sysMenu := service.SysMenuService{}.QuerySysMenuById(menuId)
	c.JSON(200, response.OkWithData(sysMenu))
}

// UpdateSysMenu 更新菜单
func (e SysMenuController) UpdateSysMenu(c *gin.Context) {
	updateReq := dto.SysMenuUpdateReqDTO{}
	err := c.ShouldBindJSON(updateReq)
	if err != nil {
		errors.ThrowException(errors.ParamsError)
		return
	}
	service.SysMenuService{}.UpdateSysMenu(updateReq)
	c.JSON(200, response.Ok())
}

// GetSysMenuTree 查询菜单树
func (e SysMenuController) GetSysMenuTree(c *gin.Context) {
	service.SysMenuService{}.QuerySysMenuTree()
}

// DeleteSysMenu 删除指定菜单
func (e SysMenuController) DeleteSysMenu(c *gin.Context) {
	deleteReq := dto.SysMenuDeleteReqDTO{}
	err := c.ShouldBindJSON(deleteReq)
	if err != nil {
		errors.ThrowException(errors.ParamsError)
		return
	}
	service.SysMenuService{}.DeleteSysMenu(deleteReq.Ids)
	c.JSON(200, response.Ok())
}

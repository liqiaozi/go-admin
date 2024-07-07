package controller

import (
	"lixuefei.com/go-admin/internal/service"
)

type ControllerGroup struct {
	SysUserController
}

var ControllerGroupApp = new(ControllerGroup)

var (
	sysUserService = service.ServiceGroupApp.SystemServiceGroup.SysUserService
)

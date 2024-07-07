package controller

import (
	"lixuefei.com/go-admin/internal/service"
)

type ControllerGroup struct {
	SysUserController
	CaptchaController
}

var ControllerGroupApp = new(ControllerGroup)

var (
	sysUserService = service.ServiceGroupApp.SystemServiceGroup.SysUserService
)

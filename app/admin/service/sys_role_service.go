package service

import (
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"lixuefei.com/go-admin/app/admin/dao"
	"lixuefei.com/go-admin/app/admin/model"
	"lixuefei.com/go-admin/app/admin/model/dto"
	"lixuefei.com/go-admin/common/component/logger"
	"lixuefei.com/go-admin/common/errors"
	"lixuefei.com/go-admin/common/utils"
)

type SysRoleService struct{}

func (s SysRoleService) AddSysRole(sysRoleDTO *dto.SysRoleDTO) int {
	logger.Log.Infof("新增系统角色, params: %v", utils.Object2JsonString(sysRoleDTO))

	sysRole := model.SysRole{}
	copier.Copy(sysRole, &sysRoleDTO)
	err, roleId := dao.SysRoleDao{}.AddSysRole(&sysRole)
	if err != nil {
		logger.Log.Error("新增系统角色异常", zap.Error(err))
		errors.ThrowException(errors.SysRoleAddError)
	}
	return roleId
}

func (s SysRoleService) QuerySysRoleById(roleId int) *model.SysRole {
	logger.Log.Infof("查询系统角色详情, params: %v", roleId)

	sysRole, err := dao.SysRoleDao{}.QuerySysRoleById(roleId)
	if err != nil {
		logger.Log.Errorf("查询角色详情异常，%v", err.Error())
		errors.ThrowExceptionWithMsg(errors.SysRoleCommonError, "查询角色详情异常")
	}
	menuIds, err := dao.SysRoleDao{}.QueryRoleMenuId(roleId)
	if err != nil {
		logger.Log.Errorf("获取角色菜单异常, %v", err.Error())
		errors.ThrowExceptionWithMsg(errors.SysRoleCommonError, "获取角色菜单异常")
	}
	sysRole.MenuIds = menuIds
	return sysRole
}

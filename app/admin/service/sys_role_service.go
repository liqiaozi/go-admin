package service

import (
	"github.com/jinzhu/copier"
	"lixuefei.com/go-admin/app/admin/dao"
	"lixuefei.com/go-admin/app/admin/model"
	"lixuefei.com/go-admin/app/admin/model/dto"
	"lixuefei.com/go-admin/common/component/logger"
	"lixuefei.com/go-admin/common/errors"
	"lixuefei.com/go-admin/common/response"
	"lixuefei.com/go-admin/common/utils"
)

type SysRoleService struct{}

func (s SysRoleService) AddSysRole(sysRoleDTO *dto.SysRoleDTO) int {
	logger.Log.Infof("新增角色, params: %v", utils.Object2JsonString(sysRoleDTO))
	sysRole := model.SysRole{}
	copier.Copy(&sysRole, &sysRoleDTO)
	err := dao.SysRoleDao{}.AddSysRole(&sysRole)
	if err != nil {
		logger.Log.Error("新增角色发生异常, err:", err.Error())
		errors.ThrowException(errors.SysRoleAddError)
	}
	return sysRole.RoleId
}

func (s SysRoleService) QuerySysRoleById(roleId int) *model.SysRole {
	logger.Log.Infof("查询角色详情, params: %v", roleId)

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

func (s SysRoleService) UpdateSysRole(updateRoleDTO *dto.SysRoleDTO) {
	logger.Log.Infof("更新角色, params: %v", utils.Object2JsonString(updateRoleDTO))
	roleList, err := dao.SysRoleDao{}.QuerySysRoleByRoleKeyOrRoleName(updateRoleDTO.RoleKey, updateRoleDTO.RoleName)
	if err != nil {
		logger.Log.Errorf("更新角色异常, %v", err.Error())
		errors.ThrowExceptionWithMsg(errors.SysRoleUpdateError, "更新角色异常")
	}
	for _, role := range roleList {
		if role.RoleId != updateRoleDTO.RoleId {
			logger.Log.Errorf("更新角色异常, 角色编码或角色名称已被占用")
			errors.ThrowExceptionWithMsg(errors.SysRoleExistError, "角色编码或角色名称已被占用")
		}
	}

	sysRole := model.SysRole{}
	copier.Copy(&sysRole, &updateRoleDTO)
	err = dao.SysRoleDao{}.UpdateSysRole(&sysRole)
	if err != nil {
		logger.Log.Error("更新角色异常", err.Error())
		errors.ThrowExceptionWithMsg(errors.SysRoleUpdateError, "更新角色异常")
	}
}

func (s SysRoleService) QuerySysRoleListByPage(rolePageQuery *dto.SysRolePageQueryDTO) *response.PageResult {
	logger.Log.Infof("分页查询角色列表, params: %v", utils.Object2JsonString(rolePageQuery))
	list, total, err := dao.SysRoleDao{}.QuerySysRoleByPage(rolePageQuery.Page, rolePageQuery.PageSize, rolePageQuery.RoleKey, rolePageQuery.RoleName, rolePageQuery.Status)
	if err != nil {
		logger.Log.Error("分页查询角色异常", err.Error())
		errors.ThrowExceptionWithMsg(errors.SysRoleUpdateError, "分页查询角色异常")
	}
	return &response.PageResult{
		Page:     rolePageQuery.Page,
		PageSize: rolePageQuery.PageSize,
		Total:    total,
		List:     list,
	}
}

func (s SysRoleService) DeleteSysRoleById(roleId int) {
	logger.Log.Infof("删除角色, roleId: %v", roleId)
	err := dao.SysRoleDao{}.DeleteSysRoleByRoleId(roleId)
	if err != nil {
		logger.Log.Error("删除角色异常", err.Error())
		errors.ThrowExceptionWithMsg(errors.SysRoleDeleteError, "删除角色异常")
	}
}

func (s SysRoleService) UpdateSysRoleStatus(roleId int, status string) {
	logger.Log.Infof("更新角色状态, roleId:%v, status: %v", roleId, status)
	sysRole, err := dao.SysRoleDao{}.QuerySysRoleById(roleId)
	if err != nil {
		logger.Log.Errorf("查询角色详情异常，%v", err.Error())
		errors.ThrowExceptionWithMsg(errors.SysRoleCommonError, "查询角色详情异常")
	}
	sysRole.Status = status
	err = dao.SysRoleDao{}.UpdateSysRole(sysRole)
	if err != nil {
		logger.Log.Error("更新角色状态异常", err.Error())
		errors.ThrowExceptionWithMsg(errors.SysRoleUpdateError, "更新角色状态异常")
	}
}

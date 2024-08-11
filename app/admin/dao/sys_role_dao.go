package dao

import (
	"errors"
	"gorm.io/gorm"
	"lixuefei.com/go-admin/app/admin/model"
	"lixuefei.com/go-admin/global"
	"time"
)

type SysRoleDao struct{}

func (d SysRoleDao) AddSysRole(sysRole *model.SysRole) (error, int) {
	sysRole.CreatedTime = time.Now().UnixMilli()
	sysRole.UpdatedTime = time.Now().UnixMilli()
	// rolekey和roleName不能重复
	if errors.Is(global.App.DB.Where("c_role_key = ? or c_role_name = ?", sysRole.RoleKey, sysRole.RoleName).First(&model.SysRole{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复的角色编码或角色名称，请进行修改"), sysRole.RoleId
	}
	return global.App.DB.Create(sysRole).Error, sysRole.RoleId
}

func (d SysRoleDao) QuerySysRoleById(roleId int) (*model.SysRole, error) {
	sysRole := model.SysRole{}
	err := global.App.DB.Where("i_role_id = ?", roleId).First(&sysRole).Error
	return &sysRole, err
}

// QueryRoleMenuId 获取角色对应的菜单ids
func (d SysRoleDao) QueryRoleMenuId(roleId int) ([]int, error) {
	menuIds := make([]int, 0)
	model := model.SysRole{}
	model.RoleId = roleId
	if err := global.App.DB.Model(&model).Preload("SysMenus").First(&model).Error; err != nil {
		return nil, err
	}
	l := *model.SysMenus
	for i := 0; i < len(l); i++ {
		menuIds = append(menuIds, l[i].MenuId)
	}
	return menuIds, nil
}

package dao

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"lixuefei.com/go-admin/app/admin/model"
	"lixuefei.com/go-admin/global"
	"strings"
	"time"
)

type SysRoleDao struct{}

// AddSysRole 新增角色
func (d SysRoleDao) AddSysRole(sysRole *model.SysRole) error {
	var err error
	sysRole.CreatedTime = time.Now().UnixMilli()
	sysRole.UpdatedTime = time.Now().UnixMilli()
	err = global.App.DB.Where("c_role_key = ? or c_role_name = ?", sysRole.RoleKey, sysRole.RoleName).First(&model.SysRole{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return global.App.DB.Create(sysRole).Error
	}
	return err
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

// QuerySysRoleByRoleKeyOrRoleName 根据roleKey和roleName查询角色列表
func (d SysRoleDao) QuerySysRoleByRoleKeyOrRoleName(roleKey string, roleName string) ([]*model.SysRole, error) {
	var sysRoles []*model.SysRole
	tx := global.App.DB.Where("c_role_key = ? or c_role_name = ?", roleKey, roleName).Find(&sysRoles)
	if err := tx.Error; err != nil {
		return sysRoles, err
	}
	return sysRoles, nil
}

// UpdateSysRole 更新角色
func (d SysRoleDao) UpdateSysRole(updateRole *model.SysRole) error {
	return global.App.DB.Save(updateRole).Error
}

// QuerySysRoleByPage 分页查询角色列表
func (d SysRoleDao) QuerySysRoleByPage(pageNo int, pageSize int, key string, name string, status string) ([]model.SysRole, int64, error) {
	var results []model.SysRole
	var total int64

	tx := global.App.DB.Model(&model.SysRole{}).Count(&total).Offset(pageNo - 1).Limit(pageSize)
	if strings.EqualFold(key, "") {
		tx.Where("c_role_key = ?", key)
	}
	if strings.EqualFold(name, "") {
		tx.Where("c_role_name like ?", "%"+key+"%")
	}
	if strings.EqualFold(status, "") {
		tx.Where("c_status = ?", status)
	}
	err := tx.Find(&results).Error
	if err != nil {
		return results, total, err
	}
	return results, total, nil
}

func (d SysRoleDao) DeleteSysRoleByRoleId(roleId int) error {
	var model = model.SysRole{}
	tx := global.App.DB.Preload("SysMenus").First(&model, roleId)
	//删除 SysRole 时，同时删除角色所有 关联其它表 记录 (SysMenu 和 SysMenu)
	return tx.Select(clause.Associations).Delete(&model).Error
}

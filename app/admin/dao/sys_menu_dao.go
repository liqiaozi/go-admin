package dao

import (
	"errors"
	"gorm.io/gorm"
	"lixuefei.com/go-admin/app/admin/model"
	"lixuefei.com/go-admin/global"
	"time"
)

type SysMenuDao struct{}

// AddSysMenu 新增菜单
func (d SysMenuDao) AddSysMenu(sysMenu *model.SysMenu) error {
	sysMenu.CreateTime = time.Now().UnixMilli()
	sysMenu.UpdateTime = time.Now().UnixMilli()
	// 菜单名称不能重复
	err := global.App.DB.Where("c_menu_name = ?", sysMenu.MenuName).First(&model.SysMenu{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return global.App.DB.Create(sysMenu).Error
	}
	return errors.New("存在重复的菜单名称")
}

// QuerySysMenuById 根据菜单ID查询详情
func (d SysMenuDao) QuerySysMenuById(menuId int) (*model.SysMenu, error) {
	sysMenu := model.SysMenu{}
	err := global.App.DB.Where("menu_id = ?", menuId).First(&sysMenu).Error
	return &sysMenu, err
}

func (d SysMenuDao) QueryAllMenu() ([]model.SysMenu, error) {
	var allMenus []model.SysMenu
	err := global.App.DB.Find(&allMenus).Error
	return allMenus, err
}

func (d SysMenuDao) UpdateSysMenu(sysMenu *model.SysMenu) {
	sysMenu.UpdateTime = time.Now().UnixMilli()
	global.App.DB.Save(sysMenu)
}

func (d SysMenuDao) QuerySysMenuLikePaths(paths string) []*model.SysMenu {
	menus := []*model.SysMenu{}
	global.App.DB.Where("c_paths like ?", paths).Find(&menus)
	return menus
}

func (d SysMenuDao) DeleteByIds(ids []int) error {
	var data model.SysMenu
	tx := global.App.DB.Delete(&data, ids)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

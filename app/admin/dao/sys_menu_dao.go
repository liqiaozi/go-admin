package dao

import (
	"errors"
	"gorm.io/gorm"
	"lixuefei.com/go-admin/app/admin/model"
	"lixuefei.com/go-admin/common/component/logger"
	"lixuefei.com/go-admin/common/utils"
	"lixuefei.com/go-admin/global"
	"time"
)

type SysMenuDao struct{}

// AddSysMenu 新增菜单
func (d SysMenuDao) AddSysMenu(sysMenu *model.SysMenu) error {
	var err error
	tx := global.App.DB.Debug().Begin()
	defer func() {
		if r := recover(); r != nil {
			logger.Log.Errorf("Panic recovered: %v", r)
			err = errors.New("internal server error") // 设置错误以便后续处理
			tx.Rollback()
		}
		if err != nil {
			logger.Log.Errorf("db err:%v", err)
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// 检查菜单名称不能重复
	result := global.App.DB.Where("c_menu_name = ?", sysMenu.MenuName).First(&model.SysMenu{})
	if result.RowsAffected == 0 {
		err = result.Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		} else {
			err = nil
		}
	} else {
		err = errors.New("菜单名称重复")
		return err
	}

	// 新增菜单
	sysMenu.CreateTime = time.Now().UnixMilli()
	sysMenu.UpdateTime = time.Now().UnixMilli()
	err = global.App.DB.Create(sysMenu).Error
	if err != nil {
		return err
	}

	// 处理菜单路径
	err = d.initPaths(tx, sysMenu)
	if err != nil {
		return err
	}
	return err
}

func (d SysMenuDao) initPaths(tx *gorm.DB, sysMenu *model.SysMenu) error {
	var err error
	var data model.SysMenu
	parentMenu := new(model.SysMenu)

	if sysMenu.ParentId != 0 {
		err = tx.Model(&data).First(parentMenu, sysMenu.ParentId).Error
		if err != nil {
			return err
		}
		if parentMenu.Paths == "" {
			err = errors.New("父级paths异常")
			return err
		}
		sysMenu.Paths = parentMenu.Paths + "/" + utils.IntToString(sysMenu.MenuId)
	} else {
		sysMenu.Paths = "/0/" + utils.IntToString(sysMenu.MenuId)
	}

	// 更新菜单路径
	err = tx.Model(&data).Where("i_menu_id = ?", sysMenu.MenuId).Update("c_paths", sysMenu.Paths).Error
	return err
}

// QuerySysMenuById 根据菜单ID查询详情
func (d SysMenuDao) QuerySysMenuById(menuId int) (*model.SysMenu, error) {
	var sysMenu model.SysMenu
	err := global.App.DB.Where("i_menu_id = ?", menuId).First(&sysMenu).Error
	if err != nil {
		return nil, err
	}
	return &sysMenu, nil
}

func (d SysMenuDao) QueryAllMenu() ([]model.SysMenu, error) {
	var allMenus []model.SysMenu
	err := global.App.DB.Find(&allMenus).Error
	if err != nil {
		return nil, err
	}
	return allMenus, nil
}

func (d SysMenuDao) UpdateSysMenu(sysMenu *model.SysMenu) {
	sysMenu.UpdateTime = time.Now().UnixMilli()
	global.App.DB.Save(sysMenu)
}

func (d SysMenuDao) QuerySysMenuLikePaths(paths string) []*model.SysMenu {
	var menus []*model.SysMenu
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

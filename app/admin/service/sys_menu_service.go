package service

import (
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"lixuefei.com/go-admin/app/admin/dao"
	"lixuefei.com/go-admin/app/admin/model"
	"lixuefei.com/go-admin/app/admin/model/dto"
	"lixuefei.com/go-admin/common/component/logger"
	"lixuefei.com/go-admin/common/errors"
	"strings"
)

type SysMenuService struct{}

// AddSysMenu 新增菜单
func (s SysMenuService) AddSysMenu(sysMenuAddReqDTO *dto.SysMenuAddReqDTO) int {
	sysMenu := model.SysMenu{}
	copier.Copy(sysMenu, sysMenuAddReqDTO)

	err, menuId := dao.SysMenuDao{}.AddSysMenu(&sysMenu)
	if err != nil {
		logger.Log.Error("新增菜单失败", zap.Error(err))
		errors.ThrowException(errors.SysMenuAddError)
	}
	return menuId
}

// QuerySysMenuById 根据菜单ID查询详情
func (s SysMenuService) QuerySysMenuById(menuId int) *model.SysMenu {
	sysMenu, err := dao.SysMenuDao{}.QuerySysMenuById(menuId)
	if err != nil {
		logger.Log.Error("查询菜单失败，", zap.Error(err))
		errors.ThrowException(errors.SysMenuNotFoundError)
	}
	return sysMenu
}

func (s SysMenuService) QuerySysMenuTree() {
	// 查询出所有的菜单列表信息
	allMenus, err := dao.SysMenuDao{}.QueryAllMenu()
	if err != nil {
		logger.Log.Error("查询菜单失败，", zap.Error(err))
		errors.ThrowExceptionWithMsg(errors.SysMenuCommonError, "查询菜单列表异常")
	}
	// 根据父亲id进行分组
	treeMap := make(map[int][]model.SysMenu)
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	// 将子菜单放到父菜单的children中
	menus := treeMap[0]
	for i := 0; i < len(menus); i++ {
		getBaseChildrenList(&menus[i], treeMap)
	}
}

// 获取菜单的子菜单
func getBaseChildrenList(menu *model.SysMenu, treeMap map[int][]model.SysMenu) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		getBaseChildrenList(&menu.Children[i], treeMap)
	}
}

func (s SysMenuService) UpdateSysMenu(req dto.SysMenuUpdateReqDTO) {
	// 查询菜单是否存在
	sysMenu, err := dao.SysMenuDao{}.QuerySysMenuById(req.MenuId)
	if err != nil {
		logger.Log.Error("查询菜单失败，", zap.Error(err))
		errors.ThrowException(errors.SysMenuNotFoundError)
	}
	// 这里需要进行事务处理
	// 更新菜单自身信息
	newSysMenu := model.SysMenu{}
	copier.Copy(&newSysMenu, req)
	dao.SysMenuDao{}.UpdateSysMenu(&newSysMenu)
	// 更新本菜单的子菜单关联的路径信息
	oldPaths := sysMenu.Paths
	menuList := dao.SysMenuDao{}.QuerySysMenuLikePaths(oldPaths)
	for _, v := range menuList {
		v.Paths = strings.Replace(v.Paths, oldPaths, req.Paths, 1)
		dao.SysMenuDao{}.UpdateSysMenu(v)
	}
}

func (s SysMenuService) DeleteSysMenu(ids []int) {
	err := dao.SysMenuDao{}.DeleteByIds(ids)
	if err != nil {
		logger.Log.Error("删除菜单发生异常", zap.Error(err))
	}
}

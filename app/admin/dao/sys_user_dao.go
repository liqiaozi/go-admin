package dao

import (
	"lixuefei.com/go-admin/app/admin/model"
	"lixuefei.com/go-admin/common/utils/stringutils"
	"lixuefei.com/go-admin/global"
)

type SysUserDao struct{}

func (s SysUserDao) QueryUserList(pageNo int, pageSize int, keyword string) (list []model.SysUser, count int64, err error) {
	var userList []model.SysUser
	tx := global.App.DB.Model(&model.SysUser{})
	if !stringutils.IsEmpty(keyword) {
		tx.Where("c_username like ?", "%"+keyword+"%")
	}
	var total int64
	err = tx.Count(&total).Error
	if err != nil {
		return
	}

	err = tx.Limit(pageSize).Offset(pageSize * (pageNo - 1)).Find(&userList).Error
	return userList, total, err
}

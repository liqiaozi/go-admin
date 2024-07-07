package dao

import (
	"lixuefei.com/go-admin/common/utils/stringutils"
	"lixuefei.com/go-admin/global"
	"lixuefei.com/go-admin/internal/models"
)

type SysUserDao struct{}

func (s SysUserDao) QueryUserList(pageNo int, pageSize int, keyword string) (list []models.SysUserEntity, count int64, err error) {
	var userList []models.SysUserEntity
	tx := global.App.DB.Model(&models.SysUserEntity{})
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

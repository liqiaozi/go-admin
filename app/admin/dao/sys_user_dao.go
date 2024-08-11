package dao

import (
	"errors"
	"gorm.io/gorm"
	"lixuefei.com/go-admin/app/admin/model"
	"lixuefei.com/go-admin/common/utils/stringutils"
	"lixuefei.com/go-admin/global"
	"time"
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

// QueryByUsername 根据用户名查询
func (s SysUserDao) QueryByUsername(username string) (*model.SysUser, error) {
	var user model.SysUser
	err := global.App.DB.Where("c_username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Save 新增用户
func (s SysUserDao) Save(user *model.SysUser) (*model.SysUser, error) {
	user.CreatedTime = time.Now()
	user.UpdatedTime = time.Now()
	err := global.App.DB.Create(user).Error
	return user, err
}

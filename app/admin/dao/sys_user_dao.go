package dao

import (
	"errors"
	"gorm.io/gorm"
	"lixuefei.com/go-admin/app/admin/model"
	"lixuefei.com/go-admin/app/admin/model/dto"
	"lixuefei.com/go-admin/common/utils/stringutils"
	"lixuefei.com/go-admin/global"
	"time"
)

type SysUserDao struct{}

func (s SysUserDao) QueryUserListByCondition(pageInfo *dto.SysUserPageQueryDTO) (list []model.SysUser, count int64, err error) {
	var userList []model.SysUser
	tx := global.App.DB.Model(&model.SysUser{})
	if !stringutils.IsEmpty(pageInfo.Username) {
		tx.Where("c_username like ?", "%"+pageInfo.Username+"%")
	}
	if !stringutils.IsEmpty(pageInfo.NickName) {
		tx.Where("c_nickname like ?", "%"+pageInfo.NickName+"%")
	}
	if !stringutils.IsEmpty(pageInfo.Phone) {
		tx.Where("c_phone like ?", "%"+pageInfo.Phone+"%")
	}
	if !stringutils.IsEmpty(pageInfo.Sex) {
		tx.Where("c_sex = ?", pageInfo.Sex)
	}
	if !stringutils.IsEmpty(pageInfo.Status) {
		tx.Where("c_status = ?", pageInfo.Status)
	}

	var total int64
	err = tx.Count(&total).Error
	if err != nil {
		return
	}

	err = tx.Limit(pageInfo.PageSize).Offset(pageInfo.PageSize * (pageInfo.Page - 1)).Find(&userList).Error
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

// SaveUser 新增用户
func (s SysUserDao) SaveUser(user *model.SysUser) error {
	user.CreatedTime = time.Now()
	user.UpdatedTime = time.Now()
	err := global.App.DB.Create(user).Error
	return err
}

func (s SysUserDao) QueryUserById(userId int) (*model.SysUser, error) {
	var user model.SysUser
	err := global.App.DB.Where("i_user_id = ?", userId).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s SysUserDao) UpdateUser(user *model.SysUser) error {
	user.UpdatedTime = time.Now()
	err := global.App.DB.Save(user).Error
	return err
}

func (s SysUserDao) DeleteUserById(userId int) (error, int64) {
	var user model.SysUser
	tx := global.App.DB.Delete(&user, userId)
	if err := tx.Error; err != nil {
		return err, 0
	}
	return nil, tx.RowsAffected
}

func (s SysUserDao) UpdateSysUserStatus(userId int, status string) error {
	var user model.SysUser
	tx := global.App.DB.Where("i_user_id = ?", userId).First(&user)
	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return errors.New("用户不存在")

	}

	err := global.App.DB.Table(user.TableName()).Where("i_user_id = ?", userId).Update("c_status = ?", status).Error
	return err
}

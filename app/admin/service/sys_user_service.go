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

type SysUserService struct {
}

// AddSysUser 新增用户
func (s SysUserService) AddSysUser(registerUser *dto.SysUserRegisterDTO) int {
	// 检查用户名是否已经存在
	sysUser, err := dao.SysUserDao{}.QueryByUsername(registerUser.Username)
	if err != nil {
		logger.Log.Error("根据用户名查询用户信息异常: " + err.Error())
		errors.ThrowExceptionWithMsg(errors.DatabaseSqlError, "新建系统用户异常")
	}
	if sysUser != nil {
		logger.Log.Error("用户名已占用")
		errors.ThrowException(errors.SysUsernameExistError)
	}

	// 组装用户信息
	user := model.SysUser{}
	copier.Copy(user, registerUser)
	user.Password = utils.BcryptHash(registerUser.Password)
	err = dao.SysUserDao{}.SaveUser(&user)
	if err != nil {
		logger.Log.Error("创建系统用户异常： " + err.Error())
		errors.ThrowException(errors.SysUserCreateError)
	}
	return user.UserId
}

func (s SysUserService) QuerySysUserById(userId int) *model.SysUser {
	logger.Log.Infof("根据用户id查询用户信息: %d", userId)

	sysUser, err := dao.SysUserDao{}.QueryUserById(userId)
	if err != nil {
		logger.Log.Errorf("根据用户ID查询用户发生异常: %s", err.Error())
		errors.ThrowExceptionWithMsg(errors.SysUserQueryError, "根据用户ID查询用户发生异常")
	}
	if sysUser == nil {
		logger.Log.Errorf("用户不存在: %d", userId)
		errors.ThrowException(errors.SysUserNoExistError)
	}

	return sysUser
}

func (s SysUserService) QueryUserList(pageInfo *dto.SysUserPageQueryDTO) *response.PageResult {
	userList, total, err := dao.SysUserDao{}.QueryUserListByCondition(pageInfo)
	if err != nil {
		return nil
	}
	pageResult := response.PageResult{
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
		Total:    total,
		List:     userList,
	}
	return &pageResult
}

func (s SysUserService) UpdateSysUser(updateUserDTO *dto.SysUserUpdateDTO) int {
	logger.Log.Infof("更新用户信息: %s", utils.Object2JsonString(updateUserDTO))

	// 判断用户名是否存在
	sysUser, err := dao.SysUserDao{}.QueryByUsername(updateUserDTO.Username)
	if err != nil {
		logger.Log.Error("根据用户名查询用户信息异常: " + err.Error())
		errors.ThrowExceptionWithMsg(errors.DatabaseSqlError, "新建系统用户异常")
	}
	if sysUser != nil && sysUser.UserId != updateUserDTO.UserId {
		logger.Log.Error("用户名已占用")
		errors.ThrowException(errors.SysUsernameExistError)
	}

	// 组装用户信息
	user := model.SysUser{}
	copier.Copy(user, updateUserDTO)
	user.Password = utils.BcryptHash(updateUserDTO.Password)
	err = dao.SysUserDao{}.UpdateUser(&user)
	if err != nil {
		logger.Log.Error("更新系统用户异常： " + err.Error())
		errors.ThrowException(errors.SysUserUpdateError)
	}
	return user.UserId
}

func (s SysUserService) DeleteSysUserById(userId int) {
	logger.Log.Infof("根据用户id删除用户, userId: %d", userId)

	err, count := dao.SysUserDao{}.DeleteUserById(userId)
	if err != nil {
		logger.Log.Error("删除系统用户异常： " + err.Error())
		errors.ThrowException(errors.SysUserDeleteError)
	}
	if count == 0 {
		logger.Log.Error("不存在该用户id记录")
		errors.ThrowException(errors.SysUserNoExistError)
	}
}

func (s SysUserService) UpdateSysUserStatus(userId int, status string) int {
	logger.Log.Infof("更新用户状态，userId: %d, status: %s", userId, status)
	err := dao.SysUserDao{}.UpdateSysUserStatus(userId, status)
	if err != nil {
		logger.Log.Errorf("更新用户状态异常,%s", err.Error())
		errors.ThrowException(errors.SysUserUpdateError)
	}
	return userId
}

package service

import (
	"lixuefei.com/go-admin/app/admin/dao"
	"lixuefei.com/go-admin/app/admin/model"
	"lixuefei.com/go-admin/app/admin/model/dto"
	"lixuefei.com/go-admin/common/component/logger"
	"lixuefei.com/go-admin/common/errors"
	"lixuefei.com/go-admin/common/request"
	"lixuefei.com/go-admin/common/response"
	"lixuefei.com/go-admin/common/utils"
)

type SysUserService struct {
}

func (s SysUserService) QueryUserList(pageInfo request.PageInfo) *response.PageResult {
	pageNo := pageInfo.Page
	pageSize := pageInfo.PageSize
	keyword := pageInfo.Keyword
	userList, total, err := dao.SysUserDao{}.QueryUserList(pageNo, pageSize, keyword)
	if err != nil {
		return nil
	}

	pageResult := response.PageResult{
		Page:     pageNo,
		PageSize: pageSize,
		Total:    total,
		List:     userList,
	}
	return &pageResult
}

// Create 新增用户
func (s SysUserService) Create(register dto.SysUserRegister) uint {
	// 检查用户名是否已经存在
	sysUser, err := dao.SysUserDao{}.QueryByUsername(register.Username)
	if err != nil {
		logger.Log.Error("query user by username error: " + err.Error())
		errors.ThrowExceptionWithMsg(errors.DatabaseSqlError, "query user by username fail")
	}
	if sysUser != nil {
		errors.ThrowException(errors.UsernameExistError)
	}

	// 组装用户信息
	var roles []model.SysRole
	for _, roleId := range register.RoleIds {
		roles = append(roles, model.SysRole{
			RoleId: roleId,
		})
	}
	user := &model.SysUser{
		Username: register.Username,
		Password: utils.BcryptHash(register.Password),
		Nickname: register.NickName,
		Phone:    register.Phone,
		Email:    register.Email,
		Avatar:   register.Avatar,
		Status:   register.Status,
		RoleId:   register.RoleId,
		Roles:    roles,
	}
	user, err = dao.SysUserDao{}.Save(user)
	if err != nil {
		logger.Log.Error("create user error: " + err.Error())
		errors.ThrowException(errors.UserCreateError)
	}
	return user.UserId
}

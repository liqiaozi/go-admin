package dto

import (
	"lixuefei.com/go-admin/common/request"
)

type SysUserRegisterDTO struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Remark   string `json:"remark"`
	RoleId   int    `json:"roleId"`
}

type SysUserUpdateDTO struct {
	SysUserRegisterDTO
}

type SysUserStatusUpdateDTO struct {
	UserId int    `json:"userId"`
	Status string `json:"status" `
}

type SysUserPageQueryDTO struct {
	request.PageInfo
	Username string `json:"username"`
	NickName string `json:"nickname"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	RoleId   int    `json:"roleId"`
}

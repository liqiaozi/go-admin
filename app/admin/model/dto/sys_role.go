package dto

import (
	"lixuefei.com/go-admin/app/admin/model"
	"lixuefei.com/go-admin/common/request"
	"time"
)

type SysRoleDTO struct {
	RoleId      int             `json:"id"`           // 角色编码
	RoleName    string          `json:"roleName"`     // 角色名称
	Status      string          `json:"status"`       // 状态 1禁用 2正常
	RoleKey     string          `json:"roleKey"`      // 角色代码
	RoleSort    int             `json:"roleSort"`     // 角色排序
	Remark      string          `json:"remark"`       // 备注
	Admin       bool            `json:"admin"`        // 是否超级管理员
	MenuIds     []int           `json:"menuIds"`      // 关联的菜单ID
	SysMenu     []model.SysMenu `json:"sysMenu"`      // 关联的菜单信息
	CreatedTime time.Time       `json:"createdTime" ` // 创建时间
	UpdatedTime time.Time       `json:"updatedTime" ` // 更新时间
}

// SysRolePageQueryDTO 角色分页查询
type SysRolePageQueryDTO struct {
	request.PageInfo
	RoleName string `json:"roleName"` // 角色名称
	RoleKey  string `json:"roleKey"`  // 角色代码
	Status   string `json:"status"`   // 状态 1禁用 2正常
}

type UpdateRoleStatusReqDTO struct {
	RoleId int    `json:"id"`     // 角色编码
	Status string `json:"status"` // 状态 1禁用 2正常
}

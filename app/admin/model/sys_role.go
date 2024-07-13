package model

import "time"

type SysRole struct {
	RoleId        uint      `gorm:"column:i_role_id;primary_key;autoIncrement" json:"roleId" `       // 角色ID
	RoleName      string    `gorm:"column:c_role_name" json:"roleName"`                              // 角色名称
	ParentId      uint      `gorm:"column:r_parent_id" json:"parentId"`                              // 父角色ID
	DefaultRouter string    `gorm:"column:c_default_router;default:dashboard" json:"defaultRouter" ` // 默认菜单(默认dashboard)
	CreatedTime   time.Time `gorm:"column:d_create_time" json:"createdTime" `                        // 创建时间
	UpdatedTime   time.Time `gorm:"column:d_update_time" json:"updatedTime" `                        // 更新时间
	Children      []SysRole `gorm:"-" json:"children"`
	Users         []SysUser `gorm:"many2many:sys_user_role;" json:"-" `
}

func (SysRole) TableName() string {
	return "sys_role"
}

package model

type SysUserRole struct {
	SysUserId uint `gorm:"column:sys_user_id"`
	SysRoleId uint `gorm:"column:sys_role_id"`
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}

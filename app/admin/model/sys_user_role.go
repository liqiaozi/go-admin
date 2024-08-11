package model

type SysUserRole struct {
	ID        uint `gorm:"column:id;primary_key;autoIncrement" json:"id" ` // 主键
	SysUserId uint `gorm:"column:sys_user_user_id"`
	SysRoleId uint `gorm:"column:sys_role_role_id"`
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}

package model

type SysUserRole struct {
	ID        int `gorm:"column:id;primary_key;autoIncrement" json:"id" ` // 主键
	SysUserId int `gorm:"column:sys_user_id"`
	SysRoleId int `gorm:"column:sys_role_id"`
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}

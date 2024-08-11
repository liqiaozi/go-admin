package model

type SysRoleMenu struct {
	ID     int `gorm:"column:id;primary_key;autoIncrement" json:"id" ` // 主键
	RoleId int `json:"role_id" gorm:"column:role_id"`
	MenuId int `json:"role_id" gorm:"column:menu_id"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}

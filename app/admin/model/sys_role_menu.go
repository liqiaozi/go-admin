package model

type SysRoleMenu struct {
	ID     int `gorm:"column:id;primary_key;autoIncrement" json:"id" ` // 主键
	RoleId int `json:"roleId" gorm:"column:role_id"`
	MenuId int `json:"menuId" gorm:"column:menu_id"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}

package model

type SysRole struct {
	RoleId      int        `gorm:"column:i_role_id;primary_key;autoIncrement" json:"roleId" ` // 角色ID
	RoleName    string     `gorm:"column:c_role_name" json:"roleName"`                        // 角色名称
	Status      string     `gorm:"column:c_status" json:"status"`                             // 状态 1禁用 2正常
	RoleKey     string     `gorm:"column:c_role_key" json:"roleKey"`                          //角色代码
	RoleSort    int        `gorm:"column:i_role_sort" json:"roleSort"`                        //角色排序
	Remark      string     `gorm:"column:c_remark" json:"remark"`                             //备注
	Admin       bool       `gorm:"column:b_admin" json:"admin"`                               // 是否超级管理员
	CreatedTime int64      `gorm:"column:d_create_time" json:"createdTime" `                  // 创建时间
	UpdatedTime int64      `gorm:"column:d_update_time" json:"updatedTime" `                  // 更新时间
	MenuIds     []int      `json:"menuIds" gorm:"-"`
	SysMenus    *[]SysMenu `json:"sysMenus" gorm:"many2many:sys_role_menu;foreignKey:RoleId;joinForeignKey:role_id;references:MenuId;joinReferences:menu_id;"`
}

func (SysRole) TableName() string {
	return "sys_role"
}

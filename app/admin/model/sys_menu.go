package model

type SysMenu struct {
	MenuId     int    `gorm:"column:i_menu_id;primaryKey;autoIncrement" json:"menuId"`
	MenuCode   string `gorm:"column:c_menu_code" json:"menuCode"`
	MenuName   string `gorm:"column:c_menu_name" json:"menuName"`
	Icon       string `gorm:"column:c_icon" json:"icon"`
	Path       string `gorm:"column:c_path" json:"path"`
	Paths      string `gorm:"column:c_paths" json:"paths"`
	MenuType   string `gorm:"column:c_menu_type" json:"menuType" `
	Action     string `gorm:"column:c_action" json:"action"`
	Permission string `gorm:"column:c_permissions" json:"permission"`
	ParentId   int    `gorm:"column:i_parent_id" json:"parentId"`
	NoCache    bool   `gorm:"column:b_no_cache" json:"noCache"`
	Breadcrumb string `gorm:"column:c_bread_crumb" json:"breadcrumb"`
	Component  string `gorm:"column:c_component" json:"component"`
	Sort       int    `gorm:"column:i_sort" json:"sort"`
	Visible    string `gorm:"column:c_visible" json:"visible"`
	IsFrame    string `gorm:"column:is_frameD" json:"isFrame"`
	CreateTime int64  `gorm:"column:d_create_time" json:"createTime"`
	UpdateTime int64  `gorm:"column:d_update_time" json:"updateTime"`

	IsSelect bool      `gorm:"-" json:"isSelect,omitempty" `
	RoleId   int       `gorm:"-" json:"isSelect,omitempty" `
	Children []SysMenu `gorm:"-" json:"children,omitempty" `
}

func (*SysMenu) TableName() string {
	return "sys_menu"
}

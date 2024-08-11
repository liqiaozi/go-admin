package dto

// SysMenuAddReqDTO 新增菜单请求
type SysMenuAddReqDTO struct {
	MenuId     int    `json:"id"`         // 编码
	MenuName   string `json:"menuName"`   //菜单name
	Title      string `json:"title"`      //显示名称
	Icon       string `json:"icon"`       //图标
	Path       string `json:"path"`       //路径
	Paths      string `json:"paths"`      //id路径
	MenuType   string `json:"menuType"`   //菜单类型
	Action     string `json:"action"`     //请求方式
	Permission string `json:"permission"` //权限编码
	ParentId   int    `json:"parentId"`   //上级菜单
	NoCache    bool   `json:"noCache"`    //是否缓存
	Breadcrumb string `json:"breadcrumb"` //是否面包屑
	Component  string `json:"component"`  //组件
	Sort       int    `json:"sort"`       //排序
	Visible    string `json:"visible"`    //是否显示
	IsFrame    string `json:"isFrame"`    //是否frame
}

// SysMenuUpdateReqDTO 更新菜单请求
type SysMenuUpdateReqDTO struct {
	SysMenuAddReqDTO
}

type SysMenuDeleteReqDTO struct {
	Ids []int `json:"ids"`
}

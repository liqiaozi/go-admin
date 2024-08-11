package model

import "time"

type SysUser struct {
	UserId      uint      `gorm:"column:i_user_id;primary_key;autoIncrement" json:"userId" ` // 主键
	Username    string    `gorm:"column:c_username" json:"username" `                        // 用户登录名
	Password    string    `gorm:"column:c_password" json:"password" `                        // 用户登录密码
	Nickname    string    `gorm:"column:c_nickname" json:"nickname" `                        // 用户昵称
	Avatar      string    `gorm:"column:c_avatar" json:"avatar"`                             // 头像
	Phone       string    `gorm:"column:c_phone" json:"phone"`                               // 用户手机号
	Email       string    `gorm:"column:c_email" json:"email"  `                             // 用户邮箱
	Status      int       `gorm:"column:i_status" json:"status" `                            // 用户状态 1正常 2冻结
	CreatedTime time.Time `gorm:"column:d_create_time" json:"createdTime" `                  // 创建时间
	UpdatedTime time.Time `gorm:"column:d_update_time" json:"updatedTime" `                  // 更新时间
	SideMode    string    `gorm:"column:c_side_mode;default:dark" json:"sideMode" `          // 用户侧边主题
	BaseColor   string    `gorm:"column:c_base_color;default:#fff" json:"baseColor" `        // 基础颜色
	RoleId      uint      `gorm:"column:c_role_id" json:"roleId" `                           // 用户角色ID
	Role        SysRole   `gorm:"foreignKey:RoleId;references:RoleId" json:"role" `
	Roles       []SysRole `gorm:"many2many:sys_user_role" json:"roles" `
}

func (SysUser) TableName() string {
	return "sys_user"
}

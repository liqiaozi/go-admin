package model

import "time"

type SysUser struct {
	UserId      int        `gorm:"column:i_user_id;primary_key;autoIncrement" json:"userId" ` // 主键
	Username    string     `gorm:"column:c_username" json:"username" `                        // 用户登录名
	Password    string     `gorm:"column:c_password" json:"password" `                        // 用户登录密码
	Nickname    string     `gorm:"column:c_nickname" json:"nickname" `                        // 用户昵称
	Phone       string     `gorm:"column:c_phone" json:"phone"`                               // 用户手机号
	Avatar      string     `gorm:"column:c_avatar" json:"avatar"`                             // 头像
	Sex         string     `gorm:"column:c_sex" json:"sex"`                                   // 性别
	Email       string     `gorm:"column:c_email" json:"email"  `                             // 用户邮箱
	Status      string     `gorm:"column:i_status" json:"status" `                            // 用户状态 1正常 2冻结
	Remark      string     `gorm:"column:c_remark" json:"remark" `                            // 备注
	CreatedTime time.Time  `gorm:"column:d_create_time" json:"createdTime" `                  // 创建时间
	UpdatedTime time.Time  `gorm:"column:d_update_time" json:"updatedTime" `                  // 更新时间
	RoleId      int        `gorm:"column:c_role_id" json:"roleId" `                           // 用户角色ID
	RoleIds     []int      `json:"roleIds" gorm:"-"`                                          // 关联角色Id
	SysRoles    *[]SysRole `json:"sysRoles" gorm:"many2many:sys_user_role;foreignKey:UserId;joinForeignKey:sys_user_id;references:RoleId;joinReferences:sys_role_id;"`
}

func (SysUser) TableName() string {
	return "sys_user"
}

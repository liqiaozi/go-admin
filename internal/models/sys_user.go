package models

import "time"

type SysUserEntity struct {
	UserId      string    `json:"userId" gorm:"column:user_id;primarykey" `                        // 主键
	Username    string    `json:"userName" gorm:"column:c_username;comment:用户登录名"`                 // 用户登录名
	Password    string    `json:"password" gorm:"column:c_password;comment:用户登录密码"`                // 用户登录密码
	NickName    string    `json:"nickName" gorm:"column:c_nickname;comment:用户昵称"`                  // 用户昵称
	Phone       string    `json:"phone"  gorm:"column:c_phone;comment:用户手机号"`                      // 用户手机号
	Email       string    `json:"email"  gorm:"column:c_email;comment:用户邮箱"`                       // 用户邮箱
	Status      int       `json:"enable" gorm:"column:i_status;default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	Avatar      string    `json:"avatar" gorm:"column:c_avatar;comment:头像"`                        // 头像
	SideMode    string    `json:"sideMode" gorm:"column:c_side_mode;default:dark;comment:用户侧边主题"`  // 用户侧边主题
	BaseColor   string    `json:"baseColor" gorm:"column:c_base_color;default:#fff;comment:基础颜色"`  // 基础颜色
	CreatedTime time.Time `json:"createdTime" gorm:"column:d_createTime;comment:创建时间"`             // 创建时间
	UpdatedTime time.Time `json:"updatedTime" gorm:"column:d_updateTime;comment:更新时间"`             // 更新时间
}

func (SysUserEntity) TableName() string {
	return "sys_user"
}

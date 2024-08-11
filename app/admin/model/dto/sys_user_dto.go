package dto

type SysUserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Status   string `json:"status" `
	RoleId   int    `json:"roleId"`
	RoleIds  []int  `json:"roleIds"`
}

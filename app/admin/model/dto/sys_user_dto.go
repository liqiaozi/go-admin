package dto

type SysUserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Status   int    `json:"status" `
	RoleId   uint   `json:"roleId"`
	RoleIds  []uint `json:"roleIds"`
}

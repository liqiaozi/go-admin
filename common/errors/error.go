package errors

// CustomError 自定义错误码
type CustomError struct {
	ErrorCode string
	ErrorMsg  string
}

// 错误码
var (
	OK               = CustomError{"0", "success"}
	SystemError      = CustomError{"10000", "system error"}
	BusinessError    = CustomError{"10001", "业务错误"}
	ValidateError    = CustomError{"10002", "请求参数错误"}
	TokenError       = CustomError{"10003", "token error"}
	ParamsError      = CustomError{"10004", "request param valid"}
	CaptchaError     = CustomError{"10005", "get captcha error"}
	DatabaseSqlError = CustomError{"10006", "database sql error"}

	UsernameExistError = CustomError{"11000", "username already exist"}
	UserCreateError    = CustomError{"11001", "create user error"}

	SysMenuCommonError   = CustomError{"12000", "menu common error"}
	SysMenuAddError      = CustomError{"12001", "menu add error"}
	SysMenuNotFoundError = CustomError{"12002", "menu not found error"}

	SysRoleCommonError = CustomError{"12000", "角色通用异常"}
	SysRoleAddError    = CustomError{"13001", "新增角色异常"}
	SysRoleUpdateError = CustomError{"13002", "更新角色异常"}
	SysRoleDeleteError = CustomError{"13003", "删除角色异常"}
)

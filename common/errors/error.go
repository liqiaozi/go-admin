package errors

// CustomError 自定义错误码
type CustomError struct {
	ErrorCode string
	ErrorMsg  string
}

// 错误码
var (
	OK            = CustomError{"0", "success"}
	SystemError   = CustomError{"10000", "system error"}
	BusinessError = CustomError{"10001", "业务错误"}
	ValidateError = CustomError{"10002", "请求参数错误"}
	TokenError    = CustomError{"10003", "token error"}
	ParamsError   = CustomError{"10004", "request param valid"}
	CaptchaError  = CustomError{"10005", "get captcha error"}
)

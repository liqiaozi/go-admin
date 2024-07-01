package response

import "lixuefei.com/go-admin/common/errors"

// DataResult 统一返回结构体
type DataResult struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Result 构造函数
func Result(code string, msg string, data interface{}) DataResult {
	return DataResult{
		code,
		msg,
		data,
	}
}

func Ok() DataResult {
	return Result(errors.OK.ErrorCode, errors.OK.ErrorMsg, nil)
}

func OkWithData(data interface{}) DataResult {
	return Result(errors.OK.ErrorCode, errors.OK.ErrorMsg, data)
}

func Fail() DataResult {
	return Result(errors.SystemError.ErrorCode, errors.SystemError.ErrorMsg, nil)
}

func FailByError(error errors.CustomError) DataResult {
	return DataResult{
		error.ErrorCode,
		error.ErrorMsg,
		nil,
	}
}

func FailByErrorWithMsg(error errors.CustomError, message string) DataResult {
	return DataResult{
		error.ErrorCode,
		message,
		nil,
	}
}

package errors

func ThrowException(customError CustomError) {
	panic(customError)
}

func ThrowExceptionWithMsg(customError CustomError, msg string) {
	customError.ErrorMsg = msg
	panic(customError)
}

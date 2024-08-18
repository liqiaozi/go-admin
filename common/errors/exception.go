package errors

func ThrowException(customError CustomError) error {
	//panic(customError)
	panic(customError)
}

func ThrowExceptionWithMsg(customError CustomError, msg string) error {
	customError.ErrorMsg = msg
	panic(customError)
}

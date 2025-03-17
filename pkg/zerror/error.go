package zerror

import "fmt"

type ZError struct {
	errorCode uint32
	errorMsg  string
}

func (z *ZError) GetCode() uint32 {
	return z.errorCode
}

func (z *ZError) GetMsg() string {
	return z.errorMsg
}

// 实现error接口
func (z *ZError) Error() string {
	return fmt.Sprintf("Ecode: %d, Emsg: %s", z.errorCode, z.errorMsg)
}

func NewZError(code uint32, msg string) *ZError {
	return &ZError{
		code,
		msg,
	}
}

func NewZErrCode(code uint32) *ZError {
	return &ZError{
		code,
		GetErrorMsg(code),
	}
}

func NewZErrMsg(msg string) *ZError {
	return &ZError{
		ServerError,
		msg,
	}
}

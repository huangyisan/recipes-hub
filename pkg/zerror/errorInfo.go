package zerror

// error code
const (
	SUCCESS     uint32 = 200
	ServerError uint32 = 500
	DBError     uint32 = 510
	ParamsError uint32 = 511
)

// error msg
var msg map[uint32]string

func init() {
	msg = make(map[uint32]string)
	msg[SUCCESS] = "成功"
	msg[ServerError] = "服务端异常"
	msg[DBError] = "数据库异常"
	msg[ParamsError] = "请求参数异常"
}

func GetErrorMsg(errCode uint32) string {
	if _msg, ok := msg[errCode]; ok {
		return _msg
	}
	return "其他错误"
}

func IsZerrorCode(errcode uint32) bool {
	if _, ok := msg[errcode]; ok {
		return true
	}
	return false
}

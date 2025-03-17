package zresp

import (
	"fmt"
	"github.com/huangyisan/recipes-hub/pkg/zerror"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

type RespSuccess struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	// 成功的时候返回数据
	Data interface{} `json:"data"`
}

type RespError struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func successResp(data interface{}) *RespSuccess {
	return &RespSuccess{
		200,
		"Success",
		data,
	}
}

func errorResp(errCode uint32, msg string) *RespError {
	return &RespError{
		errCode,
		msg,
	}
}

func isZError(err error) (bool, *zerror.ZError) {
	causeError := errors.Cause(err)
	var te *zerror.ZError
	return errors.As(causeError, &te), te
}

func isGrpcError(err error) *status.Status {
	if gstatus, ok := status.FromError(err); ok {
		return gstatus
	}
	return nil
}

// Zresp用于封装http响应, 替换原本的http响应
func Zresp(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	// 无报错响应
	if err == nil {
		rs := successResp(resp)
		httpx.WriteJson(w, http.StatusOK, rs)
		return
	}
	// 有错误的情况
	// 定义默认错误
	errCode := zerror.ServerError
	errMsg := zerror.GetErrorMsg(errCode)
	if err != nil {
		// 检查是否为自定义的ZError
		if ok, te := isZError(err); ok {
			errCode = te.GetCode()
			errMsg = te.GetMsg()

		} else {
			// 不是自定义错误, grpc错误的情况
			gstatus := isGrpcError(err)
			if gstatus != nil {
				// 检查是否为自定义的错误类型
				grpcCode := uint32(gstatus.Code())
				if ok := zerror.IsZerrCode(grpcCode); ok {
					// 区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端,一律返回"服务端异常"
					// 所以只有当是自定义错误的时候,才需要重新赋值errMsg
					errCode = grpcCode
					errMsg = gstatus.Message()
				}
			}
		}
		// 记录日志
		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
		httpx.WriteJson(w, http.StatusOK, errorResp(errCode, errMsg))
	}

}

// 参数错误的应答
func ParamErrorResp(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s, %s", zerror.GetErrorMsg(zerror.ParamsError), err.Error())
	httpx.WriteJson(w, http.StatusOK, errorResp(zerror.ParamsError, errMsg))
}

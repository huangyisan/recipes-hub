package rpcserver

import (
	"context"
	"github.com/huangyisan/recipes-hub/pkg/zerror"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err)
		var e *zerror.ZError
		// 判断是否为自定义的error
		if errors.As(causeErr, &e) {
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
			// 转成grpc err
			err = status.Error(codes.Code(e.GetCode()), e.GetMsg())
		} else {
			// 不是自定义的错误,则直接记录,不需要修改任何东西
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
		}
	}
	return resp, err
}

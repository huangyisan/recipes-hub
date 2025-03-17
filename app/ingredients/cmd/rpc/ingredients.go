package main

import (
	"flag"
	"fmt"
	"github.com/huangyisan/recipes-hub/pkg/interceptor/rpcserver"

	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/internal/config"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/internal/server"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/internal/svc"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/ingredients.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		__.RegisterIngredientServer(grpcServer, server.NewIngredientServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 自定义日志中间件, 当有请求进入到rpc服务时候，先进入拦截器然后就是执行handler方法.
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

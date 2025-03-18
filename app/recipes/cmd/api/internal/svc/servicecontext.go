package svc

import (
	"github.com/huangyisan/recipes-hub/app/recipes/cmd/api/internal/config"
	"github.com/huangyisan/recipes-hub/app/recipes/cmd/rpc/recipe"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	RecipeRpc recipe.RecipeZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		RecipeRpc: recipe.NewRecipeZrpcClient(zrpc.MustNewClient(c.RecipeRpcConf)),
	}
}

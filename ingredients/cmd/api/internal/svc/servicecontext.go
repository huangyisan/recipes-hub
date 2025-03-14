package svc

import (
	"github.com/huangyisan/recipes-hub/ingredients/cmd/api/internal/config"
	"github.com/huangyisan/recipes-hub/ingredients/cmd/rpc/ingredient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	IngredientRpc ingredient.IngredientZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		IngredientRpc: ingredient.NewIngredientZrpcClient(zrpc.MustNewClient(c.IngredientRpcConf)),
	}
}

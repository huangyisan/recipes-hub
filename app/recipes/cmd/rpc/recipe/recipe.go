// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: recipes.proto

package recipe

import (
	"context"

	"github.com/huangyisan/recipes-hub/app/recipes/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Recipe           = __.Recipe
	RecipeCreateReq  = __.RecipeCreateReq
	RecipeCreateResp = __.RecipeCreateResp

	RecipeZrpcClient interface {
		// 菜单详情
		RecipeCreate(ctx context.Context, in *RecipeCreateReq, opts ...grpc.CallOption) (*RecipeCreateResp, error)
	}

	defaultRecipeZrpcClient struct {
		cli zrpc.Client
	}
)

func NewRecipeZrpcClient(cli zrpc.Client) RecipeZrpcClient {
	return &defaultRecipeZrpcClient{
		cli: cli,
	}
}

// 菜单详情
func (m *defaultRecipeZrpcClient) RecipeCreate(ctx context.Context, in *RecipeCreateReq, opts ...grpc.CallOption) (*RecipeCreateResp, error) {
	client := __.NewRecipeClient(m.cli.Conn())
	return client.RecipeCreate(ctx, in, opts...)
}

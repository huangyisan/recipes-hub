// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: ingredients.proto

package server

import (
	"context"

	"github.com/huangyisan/recipes-hub/ingredients/cmd/rpc/internal/logic"
	"github.com/huangyisan/recipes-hub/ingredients/cmd/rpc/internal/svc"
	"github.com/huangyisan/recipes-hub/ingredients/cmd/rpc/pb"
)

type IngredientServer struct {
	svcCtx *svc.ServiceContext
	__.UnimplementedIngredientServer
}

func NewIngredientServer(svcCtx *svc.ServiceContext) *IngredientServer {
	return &IngredientServer{
		svcCtx: svcCtx,
	}
}

// 食材详情
func (s *IngredientServer) IngredientDetail(ctx context.Context, in *__.IngredientDetailReq) (*__.IngredientDetailResp, error) {
	l := logic.NewIngredientDetailLogic(ctx, s.svcCtx)
	return l.IngredientDetail(in)
}

// 食材列表
func (s *IngredientServer) IngredientList(ctx context.Context, in *__.IngredientListReq) (*__.IngredientListResp, error) {
	l := logic.NewIngredientListLogic(ctx, s.svcCtx)
	return l.IngredientList(in)
}

// 添加食材
func (s *IngredientServer) IngredientCreate(ctx context.Context, in *__.IngredientCreateReq) (*__.IngredientCreateResp, error) {
	l := logic.NewIngredientCreateLogic(ctx, s.svcCtx)
	return l.IngredientCreate(in)
}

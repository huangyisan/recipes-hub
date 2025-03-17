package logic

import (
	"context"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/internal/svc"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IngredientListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIngredientListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IngredientListLogic {
	return &IngredientListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 食材列表
func (l *IngredientListLogic) IngredientList(in *__.IngredientListReq) (*__.IngredientListResp, error) {
	// todo: add your logic here and delete this line

	return &__.IngredientListResp{}, nil
}

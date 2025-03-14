package logic

import (
	"context"

	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/internal/svc"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IngredientDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIngredientDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IngredientDetailLogic {
	return &IngredientDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 食材详情
func (l *IngredientDetailLogic) IngredientDetail(in *__.IngredientDetailReq) (*__.IngredientDetailResp, error) {
	// todo: add your logic here and delete this line

	return &__.IngredientDetailResp{}, nil
}

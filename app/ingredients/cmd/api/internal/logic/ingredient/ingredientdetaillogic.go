package ingredient

import (
	"context"

	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/api/internal/svc"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IngredientDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// ingredient detail
func NewIngredientDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IngredientDetailLogic {
	return &IngredientDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IngredientDetailLogic) IngredientDetail(req *types.IngredientDetailReq) (resp *types.IngredientDetailResq, err error) {
	// todo: add your logic here and delete this line

	return
}

package ingredient

import (
	"context"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/ingredient"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

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

func (l *IngredientDetailLogic) IngredientDetail(req *types.IngredientDetailReq) (resp *types.IngredientDetailResp, err error) {
	ingredientDetailResp, err := l.svcCtx.IngredientRpc.IngredientDetail(l.ctx, &ingredient.IngredientDetailReq{
		Name: req.Name,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	resp = &types.IngredientDetailResp{}
	_ = copier.Copy(resp, ingredientDetailResp)
	return
}

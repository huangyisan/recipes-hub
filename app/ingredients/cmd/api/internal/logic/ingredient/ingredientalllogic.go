package ingredient

import (
	"context"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/api/internal/svc"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/api/internal/types"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/ingredient"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type IngredientAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// ingredient all list
func NewIngredientAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IngredientAllLogic {
	return &IngredientAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IngredientAllLogic) IngredientAll(req *types.IngredientAllReq) (resp *types.IngredientAllResp, err error) {
	// todo: add your logic here and delete this line
	allResp, err := l.svcCtx.IngredientRpc.IngredientAll(l.ctx, &ingredient.IngredientAllReq{
		OrderBy: req.OrderBy,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	ingredientList := allResp.GetList()
	resp = &types.IngredientAllResp{}
	if len(ingredientList) > 0 {
		for _, v := range ingredientList {
			resp.List = append(resp.List, types.Ingredient{
				Id:           v.Id,
				Name:         v.Name,
				ImageContent: v.ImageContent,
				Description:  v.Description,
			})
		}
	}
	return
}

package ingredient

import (
	"context"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/api/internal/svc"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/api/internal/types"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/ingredient"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type IngredientListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// ingredient list
func NewIngredientListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IngredientListLogic {
	return &IngredientListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IngredientListLogic) IngredientList(req *types.IngredientListReq) (resp *types.IngredientListResp, err error) {
	// todo: add your logic here and delete this line
	allResp, err := l.svcCtx.IngredientRpc.IngredientList(l.ctx, &ingredient.IngredientListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	ingredientList := allResp.GetList()
	resp = &types.IngredientListResp{}
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

package ingredient

import (
	"context"
	"github.com/huangyisan/recipes-hub/ingredients/cmd/api/internal/svc"
	"github.com/huangyisan/recipes-hub/ingredients/cmd/api/internal/types"
	"github.com/huangyisan/recipes-hub/ingredients/cmd/rpc/ingredient"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type IngredientCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// ingredent create
func NewIngredientCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IngredientCreateLogic {
	return &IngredientCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IngredientCreateLogic) IngredientCreate(req *types.IngredientCreateReq) (resp *types.IngredientCreateResp, err error) {
	// todo: add your logic here and delete this line
	ingredientCreateResp, err := l.svcCtx.IngredientRpc.IngredientCreate(l.ctx, &ingredient.IngredientCreateReq{
		Name:         req.Name,
		ImageContent: req.ImageContent,
		Description:  req.Description,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.IngredientCreateResp{}
	_ = copier.Copy(resp, ingredientCreateResp)
	return
}

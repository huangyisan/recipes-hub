package logic

import (
	"context"
	"github.com/huangyisan/recipes-hub/pkg/zerror"
	"github.com/pkg/errors"

	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/internal/svc"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IngredientAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIngredientAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IngredientAllLogic {
	return &IngredientAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 食材列表all
func (l *IngredientAllLogic) IngredientAll(in *__.IngredientAllReq) (*__.IngredientAllResp, error) {
	// todo: add your logic here and delete this line
	builder := l.svcCtx.IngredientsModel.SelectBuilder()
	all, err := l.svcCtx.IngredientsModel.FindAll(l.ctx, builder, "")
	if err != nil {
		return nil, errors.Wrapf(zerror.NewZErrCode(zerror.DBError), "FindAll error: %v", err)
	}
	logx.Infof("%v+\n", all)
	var ingredientList []*__.Ingredient
	if len(all) > 0 {
		for _, v := range all {
			ingredientList = append(ingredientList, &__.Ingredient{
				Id:           v.Id,
				Name:         v.IngredientName,
				ImageContent: v.IngredientImageContent,
				Description:  v.IngredientDescription,
			})
		}
		return &__.IngredientAllResp{
			List: ingredientList,
		}, nil
	}
	return &__.IngredientAllResp{}, nil
}

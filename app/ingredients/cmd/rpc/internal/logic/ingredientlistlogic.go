package logic

import (
	"context"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/internal/svc"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/pb"
	"github.com/huangyisan/recipes-hub/pkg/zerror"
	"github.com/pkg/errors"

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
	builder := l.svcCtx.IngredientsModel.SelectBuilder()
	listByIdASC, err := l.svcCtx.IngredientsModel.FindPageListByIdASC(l.ctx, builder, in.Page, in.PageSize)
	if err != nil {
		return nil, errors.Wrapf(zerror.NewZErrCode(zerror.DBError), "List error: %v", err)
	}
	var ingredientList []*__.Ingredient
	if len(listByIdASC) > 0 {
		for _, v := range listByIdASC {
			ingredientList = append(ingredientList, &__.Ingredient{
				Id:           v.Id,
				Name:         v.IngredientName,
				ImageContent: v.IngredientImageContent,
				Description:  v.IngredientDescription,
			})
		}
		return &__.IngredientListResp{
			List: ingredientList,
		}, nil
	}
	return &__.IngredientListResp{}, nil

}

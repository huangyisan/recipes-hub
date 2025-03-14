package logic

import (
	"context"
	"errors"
	"github.com/huangyisan/recipes-hub/ingredients/model"

	"github.com/huangyisan/recipes-hub/ingredients/cmd/rpc/internal/svc"
	"github.com/huangyisan/recipes-hub/ingredients/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IngredientCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIngredientCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IngredientCreateLogic {
	return &IngredientCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加食材
func (l *IngredientCreateLogic) IngredientCreate(in *__.IngredientCreateReq) (*__.IngredientCreateResp, error) {
	// todo: add your logic here and delete this line
	ingredientName, err := l.svcCtx.IngredientsModel.FindOneByIngredientName(l.ctx, in.Name)
	// 查询异常的情况
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}
	// 条目已经存在
	if ingredientName != nil {
		return nil, errors.New("已经存在的食材")
	}
	// 不存在条目,则新增
	_, err = l.svcCtx.IngredientsModel.Insert(l.ctx, &model.Ingredients{
		IngredientName: in.Name,
	})
	if err != nil {
		return nil, err
	}

	return &__.IngredientCreateResp{}, nil
}

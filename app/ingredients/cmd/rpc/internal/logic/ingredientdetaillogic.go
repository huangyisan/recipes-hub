package logic

import (
	"context"
	"github.com/huangyisan/recipes-hub/pkg/zerror"
	"github.com/pkg/errors"

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
	one, err := l.svcCtx.IngredientsModel.FindOneByIngredientName(l.ctx, in.Name)
	if err != nil {
		return nil, errors.Wrapf(zerror.NewZErrCode(zerror.DBError), "Ingredient: %s, err: %v", in.Name, err)
	}
	if one == nil {
		return nil, errors.Wrapf(zerror.NewZErrCode(zerror.DBNoRecord), "Ingredient: %s, err: %v", in.Name, zerror.GetErrorMsg(zerror.DBNoRecord))
	}
	return &__.IngredientDetailResp{
		Ingredient: &__.Ingredient{
			Id:           one.IngredientId,
			Name:         one.IngredientName,
			ImageContent: one.IngredientImageContent,
			Description:  one.IngredientDescription,
		},
	}, nil
}

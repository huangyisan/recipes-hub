package ingredient

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

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

	return
}

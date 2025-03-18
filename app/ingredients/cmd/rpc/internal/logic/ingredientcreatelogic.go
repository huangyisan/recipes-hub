package logic

import (
	"context"
	__ "github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/pb"
	"github.com/huangyisan/recipes-hub/app/ingredients/model"
	"github.com/huangyisan/recipes-hub/pkg/zerror"
	"github.com/pkg/errors"

	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/internal/svc"
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
		return nil, errors.Wrapf(zerror.NewZErrCode(zerror.DBError), "Ingredient: %s, err: %v", in.Name, err)
	}
	// 条目已经存在
	if ingredientName != nil {
		return nil, errors.Wrapf(zerror.NewZErrMsg("Exist ingredient"), "Ingredient: %s exist in db", in.Name)
	}

	//上传r2
	filePath, err := l.svcCtx.S3Handler.UploadFile(l.ctx, in.ImageName, in.ImageContent, in.ImageContentType)
	if err != nil {
		return nil, err
	}

	// 列出r2
	//err = l.svcCtx.S3Handler.ListObjectsOutput(l.ctx)
	//if err != nil {
	//	return nil, err
	//}
	// 不存在条目,则新增
	_, err = l.svcCtx.IngredientsModel.Insert(l.ctx, &model.Ingredients{
		IngredientName:         in.Name,
		IngredientImageContent: filePath,
		IngredientDescription:  in.Description,
	})
	if err != nil {
		return nil, err
	}

	return &__.IngredientCreateResp{
		Ingredient: &__.Ingredient{
			Name:         in.Name,
			ImageContent: filePath,
			Description:  in.Description,
		},
	}, nil
}

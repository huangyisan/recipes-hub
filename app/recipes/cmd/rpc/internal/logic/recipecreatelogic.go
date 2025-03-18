package logic

import (
	"context"
	"github.com/huangyisan/recipes-hub/app/recipes/model"
	"github.com/huangyisan/recipes-hub/pkg/zerror"
	"github.com/pkg/errors"

	"github.com/huangyisan/recipes-hub/app/recipes/cmd/rpc/internal/svc"
	"github.com/huangyisan/recipes-hub/app/recipes/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecipeCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecipeCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecipeCreateLogic {
	return &RecipeCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 菜单创建
func (l *RecipeCreateLogic) RecipeCreate(in *__.RecipeCreateReq) (*__.RecipeCreateResp, error) {
	recipeName, err := l.svcCtx.RecipesModel.FindOneByRecipeName(l.ctx, in.Name)
	// 查询异常的情况
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(zerror.NewZErrCode(zerror.DBError), "Recipe: %s, err: %v", in.Name, err)
	}
	// 条目已经存在
	if recipeName != nil {
		return nil, errors.Wrapf(zerror.NewZErrMsg("Exist recipe"), "Recipe: %s exist in db", in.Name)
	}
	// 上传r2
	filePath, err := l.svcCtx.S3Handler.UploadFile(l.ctx, in.ImageName, in.ImageContent, in.ImageContentType)
	if err != nil {
		return nil, errors.Wrapf(zerror.NewZErrMsg("Upload File error"), "Upload File error: %s", in.Name)
	}

	return &__.RecipeCreateResp{}, nil
}

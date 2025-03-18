package logic

import (
	"context"
	"database/sql"
	"github.com/huangyisan/recipes-hub/app/recipes/cmd/rpc/internal/svc"
	"github.com/huangyisan/recipes-hub/app/recipes/cmd/rpc/pb"
	"github.com/huangyisan/recipes-hub/app/recipes/model"
	"github.com/huangyisan/recipes-hub/pkg/zerror"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
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
	// 查询cuisine表,获取id,然后插入recipe表
	cuisineName, err := l.svcCtx.CuisineModel.FindOneByCuisineName(l.ctx, in.RecipeType)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(zerror.NewZErrCode(zerror.DBError), "Recipe: %s, err: %v", in.Name, err)
	}
	// 获取条目id 并创建菜谱
	cuisineId := cuisineName.Id
	currentDate := time.Now()
	_, err = l.svcCtx.RecipesModel.Insert(l.ctx, &model.Recipes{
		RecipeName:   in.Name,
		Instructions: sql.NullString{String: in.Instructions, Valid: false},
		CookingTime:  in.CookingTime,
		Difficulty:   in.Difficulty,
		RecipeType:   cuisineId,
		ImageContent: sql.NullString{String: filePath, Valid: false},
		CreationDate: currentDate,
		LastUpdated:  currentDate,
		LastCookedDate: sql.NullTime{
			Time:  currentDate,
			Valid: false,
		},
	})
	if err != nil {
		return nil, err
	}
	logx.Infof("Name: %s", in.Name)
	return &__.RecipeCreateResp{
		Recipe: &__.Recipe{
			Name:         in.Name,
			Instructions: in.Instructions,
			CookingTime:  in.CookingTime,
			Difficulty:   in.Difficulty,
			RecipeType:   in.RecipeType,
			ImageContent: filePath,
			//CreationDate: currentDate.String(),
			//LastUpdated:  currentDate.String(),
		},
	}, nil
}

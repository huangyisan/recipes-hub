package recipe

import (
	"context"
	"github.com/huangyisan/recipes-hub/app/recipes/cmd/rpc/recipe"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"io"
	"net/http"

	"github.com/huangyisan/recipes-hub/app/recipes/cmd/api/internal/svc"
	"github.com/huangyisan/recipes-hub/app/recipes/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecipeCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

// recipe create
func NewRecipeCreateLogic(r *http.Request, ctx context.Context, svcCtx *svc.ServiceContext) *RecipeCreateLogic {
	return &RecipeCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *RecipeCreateLogic) RecipeCreate(req *types.RecipeCreateReq) (resp *types.RecipeCreateResp, err error) {
	file, header, err := l.r.FormFile("image_content")
	if err != nil {
		return nil, errors.Wrapf(err, "uploadFile error")
	}
	defer file.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.Wrapf(err, "readFile error")
	}

	recipeCreateResp, err := l.svcCtx.RecipeRpc.RecipeCreate(l.ctx, &recipe.RecipeCreateReq{
		Name:             req.Name,
		ImageContent:     fileBytes,
		ImageName:        header.Filename,
		ImageContentType: header.Header.Get("Content-Type"),
		Instructions:     req.Instructions,
		CookingTime:      req.CookingTime,
		Difficulty:       req.Difficulty,
		RecipeType:       req.RecipeType,
		//CreationDate:     req.CreationDate,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "Insert err: %v", req.Name)
	}
	resp = &types.RecipeCreateResp{}
	err = copier.Copy(resp, recipeCreateResp)
	if err != nil {
		return nil, errors.Wrapf(err, "RecipeCreate err: %v", req.Name)
	}
	return
}

package ingredient

import (
	"context"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/api/internal/svc"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/api/internal/types"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/rpc/ingredient"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
)

type IngredientCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

// ingredent create
func NewIngredientCreateLogic(r *http.Request, ctx context.Context, svcCtx *svc.ServiceContext) *IngredientCreateLogic {
	return &IngredientCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *IngredientCreateLogic) IngredientCreate(req *types.IngredientCreateReq) (resp *types.IngredientCreateResp, err error) {
	file, header, err := l.r.FormFile("image_content")
	if err != nil {
		return nil, errors.Wrapf(err, "uploadFile error")
	}
	defer file.Close()
	// 读取文件内容
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.Wrapf(err, "readFile error")
	}

	ingredientCreateResp, err := l.svcCtx.IngredientRpc.IngredientCreate(l.ctx, &ingredient.IngredientCreateReq{
		Name:             req.Name,
		ImageContent:     fileBytes,
		ImageName:        header.Filename,
		ImageContentType: header.Header.Get("Content-Type"),
		Description:      req.Description,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req.Name)
	}
	resp = &types.IngredientCreateResp{}
	_ = copier.Copy(resp, ingredientCreateResp)
	return
}

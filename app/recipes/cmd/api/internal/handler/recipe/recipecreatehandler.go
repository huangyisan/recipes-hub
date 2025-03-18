package recipe

import (
	"github.com/huangyisan/recipes-hub/pkg/zresp"
	"net/http"

	"github.com/huangyisan/recipes-hub/app/recipes/cmd/api/internal/logic/recipe"
	"github.com/huangyisan/recipes-hub/app/recipes/cmd/api/internal/svc"
	"github.com/huangyisan/recipes-hub/app/recipes/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// recipe create
func RecipeCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecipeCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			zresp.ParamErrorResp(r, w, err)
			return
		}

		l := recipe.NewRecipeCreateLogic(r, r.Context(), svcCtx)
		resp, err := l.RecipeCreate(&req)
		zresp.Zresp(r, w, resp, err)
	}
}

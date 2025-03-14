package ingredient

import (
	"net/http"

	"github.com/huangyisan/recipes-hub/ingredients/cmd/api/internal/logic/ingredient"
	"github.com/huangyisan/recipes-hub/ingredients/cmd/api/internal/svc"
	"github.com/huangyisan/recipes-hub/ingredients/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// ingredent create
func IngredientCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IngredientCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ingredient.NewIngredientCreateLogic(r.Context(), svcCtx)
		resp, err := l.IngredientCreate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

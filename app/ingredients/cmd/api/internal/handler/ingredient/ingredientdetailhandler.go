package ingredient

import (
	"github.com/huangyisan/recipes-hub/pkg/zresp"
	"net/http"

	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/api/internal/logic/ingredient"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/api/internal/svc"
	"github.com/huangyisan/recipes-hub/app/ingredients/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// ingredient detail
func IngredientDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IngredientDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			zresp.ParamErrorResp(r, w, err)
			return
		}

		l := ingredient.NewIngredientDetailLogic(r.Context(), svcCtx)
		resp, err := l.IngredientDetail(&req)
		zresp.Zresp(r, w, resp, err)
	}
}

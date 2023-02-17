package userOpt

import (
	"net/http"

	"douyin/pkg/favorite/api/internal/logic/userOpt"
	"douyin/pkg/favorite/api/internal/svc"
	"douyin/pkg/favorite/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavoriteOptHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteOptReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := userOpt.NewFavoriteOptLogic(r.Context(), svcCtx)
		resp, err := l.FavoriteOpt(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

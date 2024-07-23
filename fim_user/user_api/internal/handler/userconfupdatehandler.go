package handler

import (
	"net/http"

	"fim/fim_user/user_api/internal/logic"
	"fim/fim_user/user_api/internal/svc"
	"fim/fim_user/user_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserConfUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserConfUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserConfUpdateLogic(r.Context(), svcCtx)
		resp, err := l.UserConfUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

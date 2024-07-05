package handler

import (
	"net/http"

	"fim/fim_auth/auth_api/internal/logic"
	"fim/fim_auth/auth_api/internal/svc"
	"fim/fim_auth/auth_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func third_party_loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewThird_party_loginLogic(r.Context(), svcCtx)
		resp, err := l.Third_party_login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

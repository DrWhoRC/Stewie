package handler

import (
	"net/http"

	"fim/fim_auth/auth_api/internal/logic"
	"fim/fim_auth/auth_api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func authenticationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAuthenticationLogic(r.Context(), svcCtx)
		resp, err := l.Authentication()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package handler

import (
	"net/http"

	"fim/fim_auth/auth_api/internal/logic"
	"fim/fim_auth/auth_api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func third_party_login_infoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewThird_party_login_infoLogic(r.Context(), svcCtx)
		resp, err := l.Third_party_login_info()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package handler

import (
	"net/http"

	"fim/fim_chat/chat_api/internal/logic"
	"fim/fim_chat/chat_api/internal/svc"
	"fim/fim_chat/chat_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatSessionDisplayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatSessionDisplayRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewChatSessionDisplayLogic(r.Context(), svcCtx)
		resp, err := l.ChatSessionDisplay(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

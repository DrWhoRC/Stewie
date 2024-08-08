package handler

import (
	"net/http"

	"fim/fim_chat/chat_api/internal/logic"
	"fim/fim_chat/chat_api/internal/svc"
	"fim/fim_chat/chat_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatHistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatHistoryRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewChatHistoryLogic(r.Context(), svcCtx)
		resp, err := l.ChatHistory(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

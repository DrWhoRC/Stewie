package handler

import (
	"fmt"
	"net/http"

	"fim/fim_chat/chat_api/internal/svc"
	"fim/fim_chat/chat_api/internal/types"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatRequest
		if err := httpx.ParseHeaders(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		var upGrader = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				//鉴权 true放行
				return true
			},
		}
		conn, err := upGrader.Upgrade(w, r, nil)
		if err != nil {
			logx.Error(err)
			httpx.Error(w, err)
			return
		}
		defer conn.Close()
		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				logx.Error(err)
				fmt.Println(err)
				break
			}
			fmt.Println(string(p))
			conn.WriteMessage(websocket.TextMessage, []byte("您好，我现在有事不在，一会再和您联系"))
		}
	}
}

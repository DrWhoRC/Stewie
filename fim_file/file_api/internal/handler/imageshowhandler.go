package handler

import (
	"net/http"
	"os"
	"path"

	"fim/fim_file/file_api/internal/svc"
	"fim/fim_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ImageShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		filepath := path.Join("uploads", req.ImageType, req.ImageName)
		byteData, err := os.ReadFile(filepath)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		w.Write(byteData)
	}
}

package handler

import (
	"errors"
	"net/http"
	"os"

	"fim/fim_file/file_api/internal/svc"
	"fim/fim_file/file_api/internal/types"
	filemodel "fim/fim_file/models"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ImageShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		var fileModel filemodel.FileModel
		err := svcCtx.DB.Take(&fileModel, "uid = ?", req.Uid).Error
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("file not found"))
			return
		}

		byteData, err := os.ReadFile(fileModel.FilePath)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		w.Write(byteData)
	}
}

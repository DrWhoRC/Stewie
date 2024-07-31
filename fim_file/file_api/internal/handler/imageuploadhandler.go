package handler

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path"

	"fim/fim_file/file_api/internal/logic"
	"fim/fim_file/file_api/internal/svc"
	"fim/fim_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ImageUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		file, fileHeader, err := r.FormFile("image")
		if err != nil {
			logx.Error(err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		imageType := r.FormValue("imageType")
		if imageType == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("imageType is empty"))
			return
		}

		bytedata, _ := io.ReadAll(file)
		fileName := fileHeader.Filename
		filePath := path.Join("uploads", imageType, fileName)
		err = os.WriteFile(filePath, bytedata, 0666)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewImageUploadLogic(r.Context(), svcCtx)
		resp, err := l.ImageUpload(&req)
		resp.Url = "/" + filePath
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

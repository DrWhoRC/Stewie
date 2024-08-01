package handler

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

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

		//file size restriction
		mSize := float64(fileHeader.Size) / float64(1024*1024)
		if mSize > svcCtx.Config.FileSize {
			httpx.Error(w, errors.New("file size larger than 2M"))
			logx.Errorf("file size larger than 2M")
			return
		}

		//file suffix whitelist
		fileNameList := strings.Split(fileHeader.Filename, ".")
		if len(fileNameList) < 2 {
			httpx.Error(w, errors.New("file suffix is empty"))
			return
		}
		suffix := fileNameList[len(fileNameList)-1]
		for k, v := range svcCtx.Config.WhiteList {
			if v == suffix {
				break
			}
			if k == len(svcCtx.Config.WhiteList)-1 {
				httpx.Error(w, errors.New("file format is not supported"))
				return
			}
		}

		//file name same check
		imageType := r.FormValue("imageType")
		if imageType == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("imageType is empty"))
			return
		}

		bytedata, err := io.ReadAll(file)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

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

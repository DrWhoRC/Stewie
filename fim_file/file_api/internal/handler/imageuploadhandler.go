package handler

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strings"

	"fim/fim_file/file_api/internal/logic"
	"fim/fim_file/file_api/internal/svc"
	"fim/fim_file/file_api/internal/types"
	utils "fim/utils/pwd"

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
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		imageType := r.FormValue("imageType")
		if imageType == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("imageType is empty"))
			return
		}

		switch imageType {
		case "avatar", "group_avatar", "chat":
		default:
			httpx.ErrorCtx(r.Context(), w, errors.New(
				"imageType can only be: avatar, group_avata, chat"))
			return
		}

		fileName := fileHeader.Filename
		filePath := path.Join("uploads", imageType, fileName)

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
		//before store the file, go read the file list, if the file name is same, calculate their md5
		//if the md5 is same, rename the new one
		dirPath := path.Join(svcCtx.Config.UploadDir, imageType)
		dir, err := os.ReadDir(dirPath)
		if err != nil {
			os.MkdirAll(dirPath, 0666)
		}
		if InDir(dir, fileName) {
			//read the former file and calculate the md5
			formerFileByte, err := os.ReadFile(filePath)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}
			CurrentFileByte, err := io.ReadAll(file)
			if utils.MD5Encode(string(formerFileByte)) == utils.MD5Encode(string(CurrentFileByte)) {
				httpx.WriteJson(w, http.StatusInternalServerError, "file name invalid, autorenaming")
				randomSuffix := rand.Int()
				fileName = fileNameList[0] + fmt.Sprintf("_%d.", randomSuffix) + suffix
				fmt.Println("file rename: ", fileName)

				fileNewPath := path.Join(svcCtx.Config.UploadDir, imageType, fileName)
				err := os.WriteFile(fileNewPath, CurrentFileByte, 0666)
				if err != nil {
					httpx.ErrorCtx(r.Context(), w, err)
					return
				}

				l := logic.NewImageUploadLogic(r.Context(), svcCtx)
				resp, err := l.ImageUpload(&req)
				resp.Url = "/" + fileNewPath
				if err != nil {
					httpx.ErrorCtx(r.Context(), w, err)
				} else {
					httpx.OkJsonCtx(r.Context(), w, resp)
				}
				return
			}
			return
		}

		fmt.Println("file name: ", fileName)
		filePath = path.Join(svcCtx.Config.UploadDir, imageType, fileName)
		bytedata, err := io.ReadAll(file)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

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

func InDir(dir []os.DirEntry, file string) bool {
	for _, v := range dir {
		if v.Name() == file {
			return true
		}
	}
	return false
}

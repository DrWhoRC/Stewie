// Code generated by goctl. DO NOT EDIT.
package types

type FileRequest struct {
	UserId uint `header:"UserId"`
}

type FileResponse struct {
	Src string `json:"src"`
}

type ImageShowRequest struct {
	Uid string `path:"uid"`
}

type ImageUploadRequest struct {
	UserId uint `header:"UserId"`
}

type ImageUploadResponse struct {
	Url string `json:"url"`
}

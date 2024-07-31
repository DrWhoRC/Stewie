package logic

import (
	"context"

	"fim/fim_file/file_api/internal/svc"
	"fim/fim_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImageUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImageUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImageUploadLogic {
	return &ImageUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImageUploadLogic) ImageUpload(req *types.ImageUploadRequest) (resp *types.ImageUploadResponse, err error) {

	resp = new(types.ImageUploadResponse)
	return
}

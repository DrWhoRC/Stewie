package logic

import (
	"context"

	"fim/fim_file/file_api/internal/svc"
	"fim/fim_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImageShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImageShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImageShowLogic {
	return &ImageShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImageShowLogic) ImageShow(req *types.ImageShowRequest) error {
	// todo: add your logic here and delete this line

	return nil
}

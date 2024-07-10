package logic

import (
	"context"

	"fim/fim_auth/auth_api/internal/svc"
	"fim/fim_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Third_party_login_infoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThird_party_login_infoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Third_party_login_infoLogic {
	return &Third_party_login_infoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Third_party_login_infoLogic) Third_party_login_info() (resp *types.ThirdPartyLoginInfoResponse, err error) {
	// todo: add your logic here and delete this line

	resp = &types.ThirdPartyLoginInfoResponse{
		Data: []types.ThirdPartyLoginInfo{}, // 初始化Data切片
	}
	for _, v := range l.svcCtx.Config.ThirdPartyLoginList {
		resp.Data = append(resp.Data, types.ThirdPartyLoginInfo{
			Name: v.Name,
			Icon: v.Icon,
			Url:  v.Url,
		})
	}

	resp.Code = 1
	resp.Msg = "Third Party Info "
	return resp, nil
}

package logic

import (
	"context"
	"errors"
	"fmt"

	"fim/fim_auth/auth_api/internal/svc"
	"fim/fim_auth/auth_api/internal/types"
	"fim/fim_auth/models"
	thirdparty "fim/utils/third_party"

	"github.com/zeromicro/go-zero/core/logx"
)

type Third_party_loginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThird_party_loginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Third_party_loginLogic {
	return &Third_party_loginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Third_party_loginLogic) Third_party_login(req *types.ThirdPartyLoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line

	switch req.Flag {
	case "qq":
		info, err := thirdparty.NewQQLogin(req.Code, thirdparty.QQConfig{
			AppID:    l.svcCtx.Config.QQ.AppID,
			AppKey:   l.svcCtx.Config.QQ.AppKey,
			Redirect: l.svcCtx.Config.QQ.Redirect,
		})
		if err != nil {
			logx.Error(err)
			return nil, errors.New("login failed")
		}

		fmt.Println(info)
		var user models.UserModel
		err = l.svcCtx.DB.Take(&user, "open_id = ?", info.OpenID).Error
		if err != nil {
			// 注册逻辑
			fmt.Println("注册服务")
		}
		// 登录逻辑
	}
	return
}

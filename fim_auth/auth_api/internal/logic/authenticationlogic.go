package logic

import (
	"context"
	"errors"
	"fmt"

	"fim/fim_auth/auth_api/internal/svc"
	"fim/fim_auth/auth_api/internal/types"
	"fim/utils/jwts"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticationLogic {
	return &AuthenticationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticationLogic) Authentication(token string) (resp *types.AuthenticationResponse, err error) {
	// todo: add your logic here and delete this line
	if token == "" {
		err = errors.New("token is empty")
		return &types.AuthenticationResponse{
			Code: 1,
			Data: "token is empty",
			Msg:  "token is empty",
		}, nil
	}

	payload, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		err = errors.New("token is invalid")
		return
	}

	_, err = l.svcCtx.Redis.Get(fmt.Sprintf("logout_%d", payload.UserID)).Result()
	if err == nil {
		err = errors.New("authentication failed")
		return
	}
	resp = &types.AuthenticationResponse{
		Code: 0,
		Msg:  "Authentication successfully",
	}

	return resp, nil
}

package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	usermodel "fim/fim_user/models"
	"fim/fim_user/user_api/internal/svc"
	"fim/fim_user/user_api/internal/types"
	"fim/fim_user/user_rpc/users"
	"fim/utils/jwts"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserConfLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserConfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserConfLogic {
	return &GetUserConfLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserConfLogic) GetUserConf(req *types.UserInfoRequest, tokenString string) (resp *types.UserConfResponse, err error) {
	// todo: add your logic here and delete this line

	claims, err := jwts.ParseToken(tokenString, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return nil, err // JWT解析失败
	}

	userIdFromToken := claims.JwtPayload.UserID // 从JWT中提取用户ID

	// 比对用户ID
	if req.UserId != userIdFromToken {
		return nil, errors.New("unauthorized access") // 用户ID不匹配
	}

	res, err := l.svcCtx.UserRpc.UserConf(context.Background(), &users.UserInfoRequest{
		UserId: uint32(req.UserId),
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	var userconf usermodel.UserConfigModel

	err = json.Unmarshal(res.Data, &userconf)
	fmt.Println(string(res.Data))
	if err != nil {
		logx.Error(err)
		return
	}

	return &types.UserConfResponse{
		UserId:             userconf.ID,
		Online:             userconf.Online,
		RecallMsg:          *userconf.RecallMsg,
		FriendOnlineNotify: userconf.FriendOnlineNotify,
		Mute:               userconf.Mute,
		SecureLink:         userconf.SecureLink,
		SavePwd:            userconf.SavePwd,
		SearchUser:         userconf.SearchUser,
		Verification:       userconf.Verification,
		VerifyQuestion: types.VerifyQuestion{
			A1: userconf.VerifyQuestion.A1,
			A2: userconf.VerifyQuestion.A2,
			A3: userconf.VerifyQuestion.A3,
			Q1: userconf.VerifyQuestion.Q1,
			Q2: userconf.VerifyQuestion.Q2,
			Q3: userconf.VerifyQuestion.Q3,
		},
	}, nil
}

package logic

import (
	"context"
	"errors"
	"fmt"

	"fim/fim_auth/auth_api/internal/svc"
	"fim/fim_auth/auth_api/internal/types"

	"fim/fim_user/user_rpc/users"
	"fim/utils/jwts"
	utils "fim/utils/pwd"

	//usermodel "fim/fim_user/models"
	authmodel "fim/fim_auth/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line

	var user authmodel.UserModel
	err = l.svcCtx.DB.Where("nick_name = ?", req.UserName).First(&user).Error
	if err != nil {
		fmt.Println("User not found")
		fmt.Println("start registration")
		res, err1 := l.svcCtx.UserRpc.CreateUser(context.Background(), &users.UserCreateRequest{
			Nickname: req.UserName,
			Password: req.Password,
			Role:     1,
			Avatar:   "",
			OpenId:   "",
		})
		if err1 != nil {
			err1 = errors.New("registration failed")
			return &types.LoginResponse{
				Code: 0,
				Data: types.LoginInfo{Token: ""},
				Msg:  "registration failed",
			}, err
		}
		token, err := jwts.GenToken(jwts.JwtPayload{
			UserID:   uint(res.UserId),
			Username: req.UserName,
			Role:     1,
		}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
		if err != nil {
			logx.Error(err)
			err = errors.New("服务内部错误")
		}
		return &types.LoginResponse{
			Code: 1,
			Data: types.LoginInfo{Token: token},
			Msg:  "registration success, you are authorized",
		}, nil
	}

	if !utils.ValidPassword(req.Password, user.Salt, user.PwdWithSalt) {
		err = errors.New("Password error")
		return
	}

	token, err := jwts.GenToken(jwts.JwtPayload{
		UserID:   user.ID,
		Username: user.NickName,
		Role:     user.Role,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		logx.Error(err)
		err = errors.New("服务内部错误")

	}

	_, err = l.svcCtx.UserRpc.UserConfUpdate(context.Background(), &users.UserConfUpdateRequest{
		Online: true,
	})
	if err != nil {
		logx.Error(err)
		return
	}
	return &types.LoginResponse{
		Code: 1,
		Data: types.LoginInfo{Token: token},
		Msg:  "Authorized token, generated successfully.",
	}, nil
}

package logic

import (
	"context"
	"encoding/json"
	"fmt"

	usermodel "fim/fim_user/models"
	"fim/fim_user/user_api/internal/svc"
	"fim/fim_user/user_api/internal/types"
	"fim/fim_user/user_rpc/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserRpc.UserInfo(context.Background(), &users.UserInfoRequest{
		UserId: uint32(req.UserId),
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	var user usermodel.UserModel
	fmt.Println(string(res.Data))
	err = json.Unmarshal(res.Data, &user)
	fmt.Println(string(res.Data))
	if err != nil {
		logx.Error(err)
		return
	}

	return &types.UserInfoResponse{
		Id:       user.ID,
		Nickname: user.NickName,
		Role:     user.Role,
		Abstract: user.Abstract,
		Avatar:   user.Avatar,
	}, nil
}

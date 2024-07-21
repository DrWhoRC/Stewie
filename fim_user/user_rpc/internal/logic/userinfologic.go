package logic

import (
	"context"
	"encoding/json"
	usermodel "fim/fim_user/models"

	"fim/fim_user/user_rpc/internal/svc"
	"fim/fim_user/user_rpc/types/user_grpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user_grpc.UserInfoRequest) (*user_grpc.UserInfoResponse, error) {
	// todo: add your logic here and delete this line

	var user usermodel.UserModel

	//预加载操作只是填充了UserModel中定义的关联字段，而不会改变变量的类型。
	//所以，即使进行了预加载操作，user变量仍然保持为UserModel类型，
	//只是其中的关联字段userConfModel被填充了相应的数据。
	err := l.svcCtx.DB.Preload("UserConfigModel").Where("ID = ?", uint(in.UserId)).First(&user).Error

	if err != nil {
		return &user_grpc.UserInfoResponse{
			Data: []byte("user ID not found"),
		}, err
	}

	bytedata, _ := json.Marshal(user)

	return &user_grpc.UserInfoResponse{
		Data: bytedata,
	}, nil
}

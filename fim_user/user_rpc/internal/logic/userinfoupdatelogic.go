package logic

import (
	"context"
	"fmt"

	usermodel "fim/fim_user/models"
	"fim/fim_user/user_rpc/internal/svc"
	"fim/fim_user/user_rpc/types/user_grpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoUpdateLogic {
	return &UserInfoUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoUpdateLogic) UserInfoUpdate(in *user_grpc.UserInfoUpdateRequest) (*user_grpc.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	var user usermodel.UserModel

	if l.svcCtx.DB == nil {
		fmt.Println("db is nil")
	}
	if l.svcCtx == nil {
		fmt.Println("svcCtx is nil")
	}

	err := l.svcCtx.DB.Where("ID = ?", in.UserId).First(&user).Error
	if err != nil {
		return &user_grpc.UserInfoResponse{
			Data: []byte("user ID not found"),
		}, err
	}
	UpdateString(in.Nickname, &user.NickName)
	UpdateInt(int8(in.Role), &user.Role)
	UpdateString(in.Abstract, &user.Abstract)
	UpdateString(in.Avatar, &user.Avatar)

	fmt.Println("in.userid:", in.UserId)
	fmt.Println("in;user:", in.Abstract, ";", user.Abstract)

	err = l.svcCtx.DB.Model(&usermodel.UserModel{}).Where("id = ?", in.UserId).UpdateColumns(&user).Error
	if err != nil {
		return &user_grpc.UserInfoResponse{
			Data: []byte("update failed"),
		}, err
	}

	return &user_grpc.UserInfoResponse{
		Data: []byte("update successfully"),
	}, nil
}

// 注意这里的内部函数对外部参数的修改问题，这里的参数是指针，所以可以修改外部参数，搭配&使用效果更佳哦
func UpdateString(val1 string, val2 *string) {
	if val1 != "" {
		*val2 = val1
	}
}
func UpdateInt(val1 int8, val2 *int8) {
	if val1 != 0 {
		*val2 = val1
	}
}

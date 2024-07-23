package logic

import (
	"context"
	"fmt"
	"reflect"

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
	err := l.svcCtx.DB.Where("ID = ?", in.UserId).First(&user).Error
	if err != nil {
		return &user_grpc.UserInfoResponse{
			Data: []byte("user ID not found"),
		}, err
	}
	Update(in.Nickname, user.NickName)
	Update(in.Role, user.Role)
	Update(in.Abstract, user.Abstract)
	Update(in.Avatar, user.Avatar)

	err = l.svcCtx.DB.Model(&usermodel.UserModel{}).Where("id = ?", in.UserId).Update("nick_name", user.NickName).Error
	if err != nil {
		return &user_grpc.UserInfoResponse{
			Data: []byte("update failed"),
		}, err
	}
	err = l.svcCtx.DB.Model(&usermodel.UserModel{}).Where("id=?", in.UserId).Update("avatar", user.Avatar).Error
	if err != nil {
		return &user_grpc.UserInfoResponse{
			Data: []byte("update failed"),
		}, err

	}
	err = l.svcCtx.DB.Model(&usermodel.UserModel{}).Where("id=?", in.UserId).Update("abstract", user.Abstract).Error
	if err != nil {
		return &user_grpc.UserInfoResponse{
			Data: []byte("update failed"),
		}, err

	}
	err = l.svcCtx.DB.Model(&usermodel.UserModel{}).Where("id=?", in.UserId).Update("role", user.Role).Error
	if err != nil {
		return &user_grpc.UserInfoResponse{
			Data: []byte("update failed"),
		}, err

	}
	fmt.Println("98797129719321", user)

	return &user_grpc.UserInfoResponse{
		Data: []byte("update successfully"),
	}, nil
}
func Update(val1 any, val2 any) {
	if reflect.TypeOf(val1) == reflect.TypeOf("sss") {
		if val1 != "" {
			val2 = val1
		}
	} else {
		if val1 != 0 {
			val2 = val1
		}
	}

}

package logic

import (
	"context"
	"encoding/json"

	usermodel "fim/fim_user/models"
	"fim/fim_user/user_rpc/internal/svc"
	"fim/fim_user/user_rpc/types/user_grpc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserConfUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserConfUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserConfUpdateLogic {
	return &UserConfUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserConfUpdateLogic) UserConfUpdate(in *user_grpc.UserConfUpdateRequest) (*user_grpc.UserInfoResponse, error) {
	// todo: add your logic here and delete this line

	var userconf usermodel.UserConfigModel
	err := l.svcCtx.DB.Where("user_id = ?", in.Userid).First(&userconf).Error
	if err != nil {
		return &user_grpc.UserInfoResponse{
			Data: []byte("user ID not found"),
		}, err
	}

	UpdateBool(in.Online, &userconf.Online)
	UpdateStringadvanced(in.Recallmsg, &userconf.RecallMsg)
	UpdateBool(in.FriendOnlineNotify, &userconf.FriendOnlineNotify)
	UpdateBool(in.Mute, &userconf.Mute)
	UpdateBool(in.SecureLink, &userconf.SecureLink)
	UpdateBool(in.SavePwd, &userconf.SavePwd)
	UpdateInt(int8(in.SearchUser), &userconf.SearchUser)
	UpdateInt(int8(in.Verification), &userconf.Verification)
	VerifyQJson, err := json.Marshal(in.VerifyQuestion)
	if err != nil {
		return &user_grpc.UserInfoResponse{
			Data: []byte("json marshal failed"),
		}, err
	}
	json.Unmarshal(VerifyQJson, &userconf.VerifyQuestion)
	err = l.svcCtx.DB.Model(&usermodel.UserConfigModel{}).Where("user_id = ?", in.Userid).UpdateColumns(&userconf).Error

	UpdateFalse("mute", userconf.Mute, int(in.Userid), l.svcCtx.DB)

	UpdateFalse("online", userconf.Online, int(in.Userid), l.svcCtx.DB)

	UpdateFalse("friend_online_notify", userconf.FriendOnlineNotify, int(in.Userid), l.svcCtx.DB)

	UpdateFalse("secure_link", userconf.SecureLink, int(in.Userid), l.svcCtx.DB)

	UpdateFalse("save_pwd", userconf.SavePwd, int(in.Userid), l.svcCtx.DB)

	if err != nil {
		return &user_grpc.UserInfoResponse{
			Data: []byte("update failed"),
		}, err
	}

	return &user_grpc.UserInfoResponse{
		Data: []byte("update successfully"),
	}, nil

}
func UpdateFalse(key string, value bool, id int, dbconn *gorm.DB) {
	if value == false {
		dbconn.Model(&usermodel.UserConfigModel{}).Where("user_id = ?", id).Select(key).Update(key, false)
	}
}

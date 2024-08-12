package logic

import (
	"context"

	usermodel "fim/fim_user/models"
	"fim/fim_user/user_rpc/internal/svc"
	"fim/fim_user/user_rpc/types/user_grpc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type IsFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFriendLogic {
	return &IsFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFriendLogic) IsFriend(in *user_grpc.IsFriendRequest) (*user_grpc.IsFriendResponse, error) {
	// todo: add your logic here and delete this line

	var friend usermodel.FriendModel
	err := l.svcCtx.DB.Model(&usermodel.FriendModel{}).
		Where("(sender_id = ? && receiver_id = ?) || (receiver_id = ? && sender_id = ?)",
			in.UserId, in.FriendId, in.UserId, in.FriendId).First(&friend).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logx.Error(err)
		return &user_grpc.IsFriendResponse{}, err
	} else if err == gorm.ErrRecordNotFound {
		return &user_grpc.IsFriendResponse{
			IsFriend: false,
		}, nil
	} else {
		return &user_grpc.IsFriendResponse{
			IsFriend: true,
		}, nil
	}
}

package logic

import (
	"context"

	"fim/fim_chat/chat_api/internal/svc"
	"fim/fim_chat/chat_api/internal/types"
	chatmodel "fim/fim_chat/models"
	"fim/fim_user/user_rpc/types/user_grpc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ChatPinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatPinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatPinLogic {
	return &ChatPinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatPinLogic) ChatPin(req *types.ChatPinRequest) (resp *types.ChatPinResponse, err error) {
	// todo: add your logic here and delete this line

	var pin chatmodel.TopUserModel
	res, err := l.svcCtx.UserRpc.IsFriend(context.Background(), &user_grpc.IsFriendRequest{
		UserId:   uint32(req.UserId),
		FriendId: uint32(req.FriendId),
	})
	if res.IsFriend == false {
		return &types.ChatPinResponse{
			Data: "You are not friends yet",
		}, nil
	}
	err = l.svcCtx.DB.Where("user_id = ? AND top_user_id = ?", req.UserId, req.FriendId).First(&pin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logx.Error(err)
		return
	}
	if pin.ID == 0 {
		err = l.svcCtx.DB.Create(&chatmodel.TopUserModel{
			UserId:    req.UserId,
			TopUserId: req.FriendId,
		}).Error
		if err != nil {
			logx.Error(err)
			return
		}
		return &types.ChatPinResponse{
			Data: "Friend Pinned Successfully",
		}, nil
	} else {
		err = l.svcCtx.DB.Delete(&chatmodel.TopUserModel{}, pin.ID).Error
		if err != nil {
			logx.Error(err)
			return
		}
		return &types.ChatPinResponse{
			Data: "Friend Unpinned Successfully",
		}, nil
	}

}

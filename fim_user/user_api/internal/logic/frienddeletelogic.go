package logic

import (
	"context"

	usermodel "fim/fim_user/models"
	"fim/fim_user/user_api/internal/svc"
	"fim/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendDeleteLogic {
	return &FriendDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendDeleteLogic) FriendDelete(req *types.FriendDeleteRequest) (resp *types.FriendDeleteResponse, err error) {
	// todo: add your logic here and delete this line

	var friendShip usermodel.FriendModel
	l.svcCtx.DB.Model(&usermodel.FriendModel{}).Where(
		"(sender_id = ? AND receiver_id = ?) OR (receiver_id = ? AND sender_id = ?)",
		req.UserId, req.FriendId, req.UserId, req.FriendId).First(&friendShip)
	if friendShip.ID == 0 {
		resp = new(types.FriendDeleteResponse)
		resp.Data = "you are not even friends"
		return
	} else {
		l.svcCtx.DB.Delete(&friendShip)
		resp = new(types.FriendDeleteResponse)
		resp.Data = "delete successfully"
	}
	return
}

package logic

import (
	"context"

	usermodel "fim/fim_user/models"
	"fim/fim_user/user_api/internal/svc"
	"fim/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FrinedNoticeUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFrinedNoticeUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FrinedNoticeUpdateLogic {
	return &FrinedNoticeUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FrinedNoticeUpdateLogic) FrinedNoticeUpdate(req *types.FriendNoticeUpdateRequest) (resp *types.FriendNoticeUpdateResponse, err error) {
	// todo: add your logic here and delete this line

	friend := usermodel.FriendModel{}
	l.svcCtx.DB.Model(&usermodel.FriendModel{}).Where(
		"(sender_id = ? AND receiver_id = ?) OR (receiver_id = ? AND sender_id = ?)",
		req.UserId, req.FriendId, req.UserId, req.FriendId).First(&friend)
	if friend.SenderID == req.UserId {
		l.svcCtx.DB.Model(&usermodel.FriendModel{}).Where("id = ?", friend.ID).Select("recv_user_notice").Update("recv_user_notice", req.Notice)
	} else {
		l.svcCtx.DB.Model(&usermodel.FriendModel{}).Where("id = ?", friend.ID).Select("send_user_notice").Update("send_user_notice", req.Notice)
	}

	return &types.FriendNoticeUpdateResponse{
		Data: "update successfully :" + req.Notice,
	}, nil
}

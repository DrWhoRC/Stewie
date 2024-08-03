package logic

import (
	"context"

	usermodel "fim/fim_user/models"
	"fim/fim_user/user_api/internal/svc"
	"fim/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendListLogic) GetFriendList(req *types.FriendListRequest) (resp *types.FriendListResponse, err error) {
	// todo: add your logic here and delete this line

	if req.Limit < 0 {
		req.Limit = 10
	}
	if req.Page < 0 {
		req.Page = 1
	}
	offset := (req.Page - 1) * req.Limit

	var friends []usermodel.FriendModel
	l.svcCtx.DB.Limit(req.Limit).Offset(offset).Find(&friends, "sender_id = ? or receiver_id = ?", req.UserId, req.UserId)

	list := []types.FriendInfoResponse{}

	for _, friend := range friends {
		var id uint
		var notice string
		if friend.SenderID == req.UserId {
			id = friend.ReceiverID
			notice = friend.RecvUserNotice
		} else {
			id = friend.SenderID
			notice = friend.SendUserNotice
		}
		var user usermodel.UserModel
		l.svcCtx.DB.Model(&usermodel.UserModel{}).Where("ID = ?", id).First(&user)
		list = append(list, types.FriendInfoResponse{
			Id:       id,
			Nickname: user.NickName,
			Role:     user.Role,
			Abstract: user.Abstract,
			Avatar:   user.Avatar,
			Notice:   notice,
		})
	}
	return &types.FriendListResponse{
		List:  list,
		Count: len(list),
	}, nil
}

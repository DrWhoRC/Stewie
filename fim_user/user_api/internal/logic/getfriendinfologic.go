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

type GetFriendInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendInfoLogic {
	return &GetFriendInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendInfoLogic) GetFriendInfo(req *types.FriendInfoRequest) (resp *types.FriendInfoResponse, err error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.UserRpc.FriendInfo(context.Background(), &users.FriendInfoRequest{
		UserId:   uint32(req.UserId),
		FriendId: uint32(req.FriendId),
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	var friend types.FriendInfoResponse
	var friend_info usermodel.FriendModel

	err = json.Unmarshal(res.Data, &friend)
	fmt.Println(string(res.Data))
	if err != nil {
		logx.Error(err)
		return
	}
	err = json.Unmarshal(res.Data, &friend_info)
	if err != nil {
		logx.Error(err)
		return
	}

	if req.UserId == friend_info.SenderID {
		return &types.FriendInfoResponse{
			Id:       friend.Id,
			Nickname: friend.Nickname,
			Role:     friend.Role,
			Abstract: friend.Abstract,
			Avatar:   friend.Avatar,
			Notice:   friend_info.RecvUserNotice,
		}, nil
	} else {
		return &types.FriendInfoResponse{
			Id:       friend.Id,
			Nickname: friend.Nickname,
			Role:     friend.Role,
			Abstract: friend.Abstract,
			Avatar:   friend.Avatar,
			Notice:   friend_info.SendUserNotice,
		}, nil
	}
}

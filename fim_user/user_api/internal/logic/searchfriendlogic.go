package logic

import (
	"context"
	"fmt"

	usermodel "fim/fim_user/models"
	"fim/fim_user/user_api/internal/svc"
	"fim/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFriendLogic {
	return &SearchFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchFriendLogic) SearchFriend(req *types.SearchRequest) (resp *types.SearchResponse, err error) {
	// todo: add your logic here and delete this line
	if req.Limit < 0 {
		req.Limit = 10
	}
	if req.Page < 0 {
		req.Page = 1
	}
	offset := (req.Page - 1) * req.Limit

	users := []usermodel.UserModel{}

	query := l.svcCtx.DB.Model(&usermodel.UserModel{}).
		Joins("JOIN user_config_models ON user_config_models.user_id = user_models.id").
		Where("user_config_models.search_user != 0").
		Select("user_models.id, user_models.nick_name, user_models.avatar, user_models.abstract").
		Limit(req.Limit).Offset(offset)

	if req.Key != "" {
		query = query.Where("user_models.nick_name LIKE ?", "%"+req.Key+"%")
	}
	err = query.Find(&users).Error
	if err != nil {
		logx.Error(err)
		return
	}
	fmt.Println("users:", users)

	friends := []usermodel.FriendModel{}
	Search_OutComes := types.SearchResponse{}

	for k, v := range users {
		l.svcCtx.DB.Model(&usermodel.FriendModel{}).Where(
			"(sender_id = ? AND receiver_id = ?) OR (receiver_id = ? AND sender_id = ?)",
			req.UserId, users[k].ID, req.UserId, users[k].ID).Find(&friends)

		ouco := types.SearchInfo{}
		if len(friends) > 0 {
			ouco.Id = v.ID
			ouco.Nickname = v.NickName
			ouco.Avatar = v.Avatar
			ouco.Abstract = v.Abstract
			ouco.IsFriend = true
		} else {
			ouco.Id = v.ID
			ouco.Nickname = v.NickName
			ouco.Avatar = v.Avatar
			ouco.Abstract = v.Abstract
			ouco.IsFriend = false
		}
		friends = []usermodel.FriendModel{}
		Search_OutComes.List = append(Search_OutComes.List, ouco)
	}
	fmt.Println(Search_OutComes)

	return &Search_OutComes, nil
}

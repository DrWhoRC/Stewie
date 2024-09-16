package logic

import (
	"context"
	"encoding/json"
	"errors"

	"fim/fim_group/group_api/internal/svc"
	"fim/fim_group/group_api/internal/types"
	groupmodel "fim/fim_group/models"
	usermodel "fim/fim_user/models"
	"fim/fim_user/user_rpc/types/user_grpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCreateLogic {
	return &GroupCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupCreateLogic) GroupCreate(req *types.GroupCreateRequest) (resp *types.GroupCreateResponse, err error) {
	// todo: add your logic here and delete this line
	var group = groupmodel.GroupModel{
		CreatorID:    req.UserId,
		IsSearchable: false,
		Verification: 2,
		Size:         1,
	}

	switch req.Mode {
	case 1: //直接创建
		if req.Name == "" {
			return nil, errors.New("Group name is empty")
		}
		group.Title = req.Name
		group.Size = int8(req.Size)
		group.IsSearchable = req.IsSearch

	case 2: //选人创建
		if len(req.UserIdList) == 0 {
			return nil, errors.New("UserIdList is empty")
		}
		var userIdList = []uint32{
			uint32(req.UserId), //先把自己加进去
		}
		for _, v := range req.UserIdList {
			userIdList = append(userIdList, uint32(v))
			res, err := l.svcCtx.UserRpc.UserInfo(context.Background(), &user_grpc.UserInfoRequest{
				UserId: uint32(v),
			})
			if err != nil {
				return nil, err
			}
			user := usermodel.UserModel{}
			err = json.Unmarshal(res.Data, &user)
			if err != nil {
				return nil, errors.New("json.Unmarshal(res.Data,&user) error")
			}
		}

	}

	return
}

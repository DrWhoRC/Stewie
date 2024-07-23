package logic

import (
	"context"

	"fim/fim_user/user_api/internal/svc"
	"fim/fim_user/user_api/internal/types"
	"fim/fim_user/user_rpc/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoUpdateLogic {
	return &UserInfoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoUpdateLogic) UserInfoUpdate(req *types.UserInfoUpdateRequest) (resp *types.UserInfoUpdateResponse, err error) {
	// todo: add your logic here and delete this line

	_, err = l.svcCtx.UserRpc.UserInfoUpdate(context.Background(), &users.UserInfoUpdateRequest{
		UserId:   uint32(req.UserId),
		Nickname: derefString(req.Nickname, ""),
		Role:     int32(derefInt(req.Role, 0)),
		Abstract: derefString(req.Abstract, ""),
		Avatar:   derefString(req.Avatar, ""),
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &types.UserInfoUpdateResponse{
		Data: resp.Data,
	}, nil
}
func derefString(s *string, defaultVal string) string {
	if s != nil {
		return *s
	}
	return defaultVal
}

// 辅助函数，安全地解引用整型指针，如果是nil则返回默认值
func derefInt(i *int8, defaultVal int8) int8 {
	if i != nil {
		return *i
	}
	return defaultVal
}

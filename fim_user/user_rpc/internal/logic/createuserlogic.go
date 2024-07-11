package logic

import (
	"context"
	"errors"
	"math/rand"
	"time"

	usermodel "fim/fim_user/models"
	utils "fim/utils/pwd"

	"fim/fim_user/user_rpc/internal/svc"
	"fim/fim_user/user_rpc/types/user_grpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user_grpc.UserCreateRequest) (*user_grpc.UserCreateResponse, error) {
	// todo: add your logic here and delete this line
	var user usermodel.UserModel
	res := l.svcCtx.DB.Where("nickName=?", in.Nickname).First(&user)
	if res.Error == nil {
		return nil, errors.New("User already exists")
	}
	user.NickName = in.Nickname
	user.Pwd = in.Password
	rand.Seed(time.Now().UnixNano())
	// 生成一个1000000到9999999之间的随机数
	user.Salt = string(rand.Intn(9000000) + 1000000)
	user.PwdWithSalt = utils.MakePassword(in.Password, user.Salt)
	user.Role = int8(in.Role)
	l.svcCtx.DB.Create(&user)
	return &user_grpc.UserCreateResponse{
		UserId: int32(user.ID),
	}, nil
}

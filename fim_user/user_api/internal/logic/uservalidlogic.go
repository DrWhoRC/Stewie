package logic

import (
	"context"

	usermodel "fim/fim_user/models"
	"fim/fim_user/user_api/internal/svc"
	"fim/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserValidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserValidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserValidLogic {
	return &UserValidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserValidLogic) UserValid(req *types.UserValidRequest) (resp *types.UserValidResponse, err error) {
	// todo: add your logic here and delete this line

	var userConf usermodel.UserConfigModel
	l.svcCtx.DB.Model(&usermodel.UserConfigModel{}).Where("user_id = ?", req.FriendId).First(&userConf)
	resp = new(types.UserValidResponse)
	resp.Verification = userConf.Verification
	switch userConf.Verification {
	case 0: //不允许添加好友
		resp.Data = "this is a private account, he/she does not allow to add friends"
	case 1: //需要验证消息
		resp.Data = "the message has been forwarded to the user, please wait for the user to confirm"
	case 2: //需要验证问题的答案
		if userConf.VerifyQuestion != nil {
			resp.VerifyQuestion = types.VerifyQuestion{
				Q1: userConf.VerifyQuestion.Q1,
				Q2: userConf.VerifyQuestion.Q2,
				Q3: userConf.VerifyQuestion.Q3,
			}
			resp.Data = "please answer questions the user set"
		}
	case 3: //需要验证问题的正确答案
		if userConf.VerifyQuestion != nil {
			resp.VerifyQuestion = types.VerifyQuestion{
				Q1: userConf.VerifyQuestion.Q1,
				Q2: userConf.VerifyQuestion.Q2,
				Q3: userConf.VerifyQuestion.Q3,
			}
			resp.Data = "please answer questions correctly the user set"
		}
	case 4: //不需要验证
		AddToFriend(req.UserId, req.FriendId, l.svcCtx.DB)
	default:
	}

	return
}
func AddToFriend(user_id_Sender, user_id_Recver uint, DB *gorm.DB) error {
	friend := usermodel.FriendModel{}
	friend.SenderID = user_id_Sender
	friend.ReceiverID = user_id_Recver

	err := DB.Create(&friend).Error
	if err != nil {
		return err
	}
	return nil
}

package logic

import (
	"context"
	"fmt"

	usermodel "fim/fim_user/models"
	"fim/fim_user/user_api/internal/svc"
	"fim/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidAnswerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewValidAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidAnswerLogic {
	return &ValidAnswerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ValidAnswerLogic) ValidAnswer(req *types.ValidAnswerRequest) (resp *types.ValidAnswerResponse, err error) {
	// todo: add your logic here and delete this line

	var userConf usermodel.UserConfigModel
	l.svcCtx.DB.Model(&usermodel.UserConfigModel{}).Where("user_id = ?", req.FriendId).First(&userConf)
	resp = new(types.ValidAnswerResponse)
	var friendVerify usermodel.FriendVerifyModel
	switch userConf.Verification {
	case 2:
		if userConf.VerifyQuestion != nil {
			friendVerify.VerifyQuestion.A1 = req.VerifyQuestion.A1
			friendVerify.VerifyQuestion.A2 = req.VerifyQuestion.A2
			friendVerify.VerifyQuestion.A3 = req.VerifyQuestion.A3
			friendVerify.VerifyQuestion.Q1 = userConf.VerifyQuestion.Q1
			friendVerify.VerifyQuestion.Q2 = userConf.VerifyQuestion.Q2
			friendVerify.VerifyQuestion.Q3 = userConf.VerifyQuestion.Q3
			err := l.svcCtx.DB.Create(&friendVerify).Error
			if err != nil {
				resp.Data = err.Error()
				return resp, err
			}
			resp.Data = "the verify message has been forwarded to the user, please wait for the user to confirm"
		}
	case 3: //需要问题的正确答案
		if userConf.VerifyQuestion != nil {
			if *req.VerifyQuestion.A1 == *userConf.VerifyQuestion.A1 && *req.VerifyQuestion.A2 == *userConf.VerifyQuestion.A2 && *req.VerifyQuestion.A3 == *userConf.VerifyQuestion.A3 {
				resp.Data = "answer correct, you have become friends with him/her"
				AddToFriend(req.UserId, req.FriendId, l.svcCtx.DB)
			} else {
				resp.Data = fmt.Sprintf("answer incorrect: %s,%s,%s,%s,%s,%s", *userConf.VerifyQuestion.A1, *userConf.VerifyQuestion.A2, *userConf.VerifyQuestion.A3, *req.VerifyQuestion.A1, *req.VerifyQuestion.A2, *req.VerifyQuestion.A3)
			}
		}
	default:
	}
	return
}

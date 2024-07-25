package logic

import (
	"context"

	"fim/fim_user/user_api/internal/svc"
	"fim/fim_user/user_api/internal/types"
	"fim/fim_user/user_rpc/types/user_grpc"
	"fim/fim_user/user_rpc/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserConfUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserConfUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserConfUpdateLogic {
	return &UserConfUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserConfUpdateLogic) UserConfUpdate(req *types.UserConfUpdateRequest) (resp *types.UserConfUpdateResponse, err error) {
	// todo: add your logic here and delete this line

	var verificationquestion user_grpc.VerifyQuestion
	verificationquestion.A1 = *req.VerifyQuestion.A1
	verificationquestion.A2 = *req.VerifyQuestion.A2
	verificationquestion.A3 = *req.VerifyQuestion.A3
	verificationquestion.Q1 = *req.VerifyQuestion.Q1
	verificationquestion.Q2 = *req.VerifyQuestion.Q2
	verificationquestion.Q3 = *req.VerifyQuestion.Q3

	rpcresp, err := l.svcCtx.UserRpc.UserConfUpdate(context.Background(), &users.UserConfUpdateRequest{
		Userid:             uint32(req.UserId),
		Online:             derefBool(req.Online, bool(false)),
		Recallmsg:          derefString(req.RecallMsg, ""),
		FriendOnlineNotify: derefBool(req.FriendOnlineNotify, bool(false)),
		Mute:               derefBool(req.Mute, bool(false)),
		SecureLink:         derefBool(req.SecureLink, bool(false)),
		SavePwd:            derefBool(req.SavePwd, bool(false)),
		SearchUser:         int32(derefInt(req.SearchUser, 0)),
		Verification:       int32(derefInt(req.Verification, 0)),
		VerifyQuestion:     &verificationquestion,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &types.UserConfUpdateResponse{
		Data: string(rpcresp.Data),
	}, nil
}

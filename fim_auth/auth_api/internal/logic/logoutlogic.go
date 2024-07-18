package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"fim/fim_auth/auth_api/internal/svc"
	"fim/fim_auth/auth_api/internal/types"
	"fim/utils/jwts"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(token string) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	if token == "" {
		err = errors.New("token is empty")
		return
	}

	claim, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		err = errors.New("token is invalid")
		return
	}
	//过期时间就是jwt失效时间，
	expiration := claim.ExpiresAt.Time.Sub(time.Now())
	//expire date is the same as JWToken

	userid_str := fmt.Sprintf("logout_%d_%s", claim.UserID, token)

	//set if not exist, 设置键值对，值为空，到了过期时间就会自动删除
	l.svcCtx.Redis.SetNX(userid_str, "", expiration)

	return &types.Response{
		Code: 0,
		Msg:  "Logout successfully",
	}, nil

}

/*防止重复使用JWT进行认证：一旦用户注销，相应的JWT虽然在过期前仍然是有效的，
但是通过在Redis中记录注销状态，应用可以拒绝基于已注销的JWT的进一步认证请求。
这是一种安全措施，以防止在用户主动注销后，其JWT被盗用。*/

/*自动清理：由于设置了过期时间，与注销的用户相关的键值对会在JWT过期时自动从Redis中删除，
这样可以避免无用数据的累积。*/

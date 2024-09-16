package redis_service

import (
	"context"
	"encoding/json"
	"errors"
	usermodel "fim/fim_user/models"
	"fim/fim_user/user_rpc/types/user_grpc"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type UserInfo struct {
	UserId   uint   `json:"userId"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
}

func GetUserBaseInfo(client *redis.Client, userRpc user_grpc.UsersClient, userID uint) (userInfo UserInfo, err error) {
	str, err := client.Get(fmt.Sprintf("user_info_%d", userID)).Result()
	if err != nil {
		// if not found in redis, get from userRpc
		res, errRpc := userRpc.UserInfo(context.Background(), &user_grpc.UserInfoRequest{
			UserId: uint32(userID),
		})
		if errRpc != nil {
			return userInfo, errors.New("redis get user info failed, and rpc get user info failed")
		}

		var user usermodel.UserModel
		err = json.Unmarshal(res.Data, &user)
		userInfo = UserInfo{
			UserId:   user.ID,
			NickName: user.NickName,
			Avatar:   user.Avatar,
		}
		byteData, _ := json.Marshal(userInfo)

		client.Set(fmt.Sprintf("user_info_%d", userID), byteData, time.Hour*2)
		return userInfo, nil
	}

	err = json.Unmarshal([]byte(str), &userInfo)
	if err != nil {
		return
	}
	return userInfo, nil
}

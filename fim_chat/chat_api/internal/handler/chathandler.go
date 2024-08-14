package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"fim/fim_chat/chat_api/internal/svc"
	"fim/fim_chat/chat_api/internal/types"
	usermodel "fim/fim_user/models"
	"fim/fim_user/user_rpc/types/user_grpc"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type UserInfo struct {
	UserId   uint   `json:"userId"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
}

// map存每个用户的ws连接
type UserWsInfo struct {
	UserInfo UserInfo
	Conn     *websocket.Conn
}

var UserWsInfoMap = make(map[uint]*UserWsInfo)
var LastRound = make(map[uint]*UserWsInfo)
var mu1 sync.Mutex
var mu2 sync.Mutex

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatRequest
		if err := httpx.ParseHeaders(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		var upGrader = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				//鉴权 true放行
				return true
			},
		}
		conn, err := upGrader.Upgrade(w, r, nil)
		if err != nil {
			logx.Error(err)
			httpx.Error(w, err)
			return
		}
		defer func() {
			conn.Close()
			delete(UserWsInfoMap, req.UserId)
		}()

		//调用户服务获取用户信息
		res, err := svcCtx.UserRpc.UserInfo(context.Background(), &user_grpc.UserInfoRequest{
			UserId: uint32(req.UserId),
		})
		if err != nil {
			logx.Error(err)
			httpx.Error(w, err)
			return
		}
		var user = usermodel.UserModel{}
		err = json.Unmarshal(res.Data, &user)
		userWsInfo := &UserWsInfo{
			UserInfo: UserInfo{
				UserId:   req.UserId,
				NickName: user.NickName,
				Avatar:   user.Avatar,
			},
			Conn: conn,
		}
		fmt.Println(user.UserConfigModel)
		UserWsInfoMap[req.UserId] = userWsInfo

		//check if he's online
		resFriendList, err := svcCtx.UserRpc.GetFriendList(context.Background(), &user_grpc.FriendListRequest{
			UserId: uint32(req.UserId),
		})
		if err != nil {
			logx.Error("Friend List:", err)
			httpx.Error(w, err)
			return
		}
		for _, friend := range resFriendList.FriendList {
			if friend.FriendOnlineNotify == true {
				if UserWsInfoMap[uint(friend.UserId)] != nil {
					//send friend online notice
					conn.WriteMessage(websocket.TextMessage,
						[]byte(fmt.Sprintf("%s is online", friend.Nickname)))
				}
			}
		}

		//friend online notice
		go func() {
			for conn != nil {
				mu1.Lock()
				for _, friend := range resFriendList.FriendList {
					time.Sleep(time.Second * 1)
					if friend.FriendOnlineNotify {
						mu2.Lock()
						if UserWsInfoMap[uint(friend.UserId)] == nil {
							fmt.Println("UserWsInfo:", UserWsInfoMap[uint(friend.UserId)])
							LastRound[uint(friend.UserId)] = nil
						}
						if UserWsInfoMap[uint(friend.UserId)] != nil {
							if LastRound[uint(friend.UserId)] == nil {

								LastRound[uint(friend.UserId)] = userWsInfo

								fmt.Println("Lastround", LastRound[uint(friend.UserId)])
								conn.WriteMessage(websocket.TextMessage,
									[]byte(fmt.Sprintf("%s just got online", friend.Nickname)))
							}
						}
						mu2.Unlock()
					}
				}
				mu1.Unlock()
			}
		}()

		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				logx.Error(err)
				fmt.Println(err)
				break
			}
			fmt.Println(string(p))
			conn.WriteMessage(websocket.TextMessage, []byte("您好，我现在有事不在，一会再和您联系"))
		}
	}
}

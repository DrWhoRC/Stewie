package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"fim/common/models/ctype"
	"fim/common/service/redis_service"
	"fim/fim_chat/chat_api/internal/svc"
	"fim/fim_chat/chat_api/internal/types"
	"fim/fim_chat/chat_rpc/types/chat_rpc"
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
		res, err := redis_service.GetUserBaseInfo(svcCtx.Redis, svcCtx.UserRpc, req.UserId)
		if err != nil {
			logx.Error(err)
			httpx.Error(w, err)
			return
		}

		userWsInfo := &UserWsInfo{
			UserInfo: UserInfo{
				UserId:   res.UserId,
				NickName: res.NickName,
				Avatar:   res.Avatar,
			},
			Conn: conn,
		}
		UserWsInfoMap[req.UserId] = userWsInfo
		fmt.Println("UserWsInfoMap:", UserWsInfoMap[req.UserId])

		//check if he's online
		resFriendList, err := svcCtx.UserRpc.GetFriendList(context.Background(), &user_grpc.FriendListRequest{
			UserId: uint32(req.UserId),
		})
		fmt.Println("resFriendList:", resFriendList)
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
		// go func() {
		// 	for conn != nil {
		// 		mu1.Lock()
		// 		for _, friend := range resFriendList.FriendList {
		// 			time.Sleep(time.Second * 1)
		// 			if friend.FriendOnlineNotify {
		// 				mu2.Lock()
		// 				if UserWsInfoMap[uint(friend.UserId)] == nil {
		// 					fmt.Println("UserWsInfo:", UserWsInfoMap[uint(friend.UserId)])
		// 					LastRound[uint(friend.UserId)] = nil
		// 				}
		// 				if UserWsInfoMap[uint(friend.UserId)] != nil {
		// 					if LastRound[uint(friend.UserId)] == nil {

		// 						LastRound[uint(friend.UserId)] = userWsInfo

		// 						fmt.Println("Lastround", LastRound[uint(friend.UserId)])
		// 						conn.WriteMessage(websocket.TextMessage,
		// 							[]byte(fmt.Sprintf("%s just got online", friend.Nickname)))
		// 					}
		// 				}
		// 				mu2.Unlock()
		// 			}
		// 		}
		// 		mu1.Unlock()
		// 	}
		// }()

		for {
			_, p, err1 := conn.ReadMessage()
			if err1 != nil {
				logx.Error(err1)
				fmt.Println("User read err", err1)
				break
			}
			var request ChatRequest
			err2 := json.Unmarshal(p, &request)
			if err2 != nil {
				logx.Error(err2)
				conn.WriteMessage(websocket.TextMessage, p)
				fmt.Println(p)
				fmt.Println("json unmarshal err2", err2)
				break
			}

			isFriendRes, err := svcCtx.UserRpc.IsFriend(context.Background(), &user_grpc.IsFriendRequest{
				UserId:   uint32(req.UserId),
				FriendId: uint32(request.ReceiverId),
			})
			if err != nil {
				logx.Error("isFriend-UserRpc err", err)
				continue
			}
			if !isFriendRes.IsFriend {
				logx.Error(err)
				conn.WriteMessage(websocket.TextMessage, []byte("You are not friends yet"))
				break
			}
			//write in db

			dbBytedata, _ := json.Marshal(request.Msg)
			svcCtx.ChatRpc.UserChat(context.Background(), &chat_rpc.UserChatRequest{
				Sender:   uint32(req.UserId),
				Receiver: uint32(request.ReceiverId),
				Msg:      dbBytedata,
			})

			SendMsg(request.ReceiverId, req.UserId, request.Msg)
		}
	}
}

type ChatRequest struct {
	ReceiverId uint      `json:"receiverId"`
	Msg        ctype.Msg `json:"msg"`
}

type ChatResponse struct {
	SenderID         uint      `json:"senderId"`
	SenderNickname   string    `json:"senderNickname"`
	SenderAvatar     string    `json:"senderAvatar"`
	ReceiverID       uint      `json:"receiverId"`
	RecerverNickname string    `json:"receiverNickname"`
	ReceiverAvatar   string    `json:"receiverAvatar"`
	CreatedAt        string    `json:"createdAt"`
	Msg              ctype.Msg `json:"msg"`
}

func SendMsg(ReceiverId uint, SenderId uint, msg ctype.Msg) {
	Receiver, ok := UserWsInfoMap[ReceiverId] //check if the receiver is online
	if !ok {
		return
	}
	Sender, ok := UserWsInfoMap[SenderId]
	if !ok {
		return
	}
	resp := ChatResponse{
		SenderID:         SenderId,
		SenderNickname:   Sender.UserInfo.NickName,
		SenderAvatar:     Sender.UserInfo.Avatar,
		ReceiverID:       ReceiverId,
		RecerverNickname: Receiver.UserInfo.NickName,
		ReceiverAvatar:   Receiver.UserInfo.Avatar,
		Msg:              msg,
		CreatedAt:        time.Now().Format("2006-01-02 15:04:05"),
	}
	byteData, err := json.Marshal(resp)
	if err != nil {
		logx.Error(err)
		fmt.Println("json marshal err", err)
		Receiver.Conn.WriteMessage(websocket.TextMessage, []byte("json marshal err"))
	}
	Receiver.Conn.WriteMessage(websocket.TextMessage, byteData)
}

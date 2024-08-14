// Code generated by goctl. DO NOT EDIT.
// Source: user_rpc.proto

package server

import (
	"context"

	"fim/fim_user/user_rpc/internal/logic"
	"fim/fim_user/user_rpc/internal/svc"
	"fim/fim_user/user_rpc/types/user_grpc"
)

type UsersServer struct {
	svcCtx *svc.ServiceContext
	user_grpc.UnimplementedUsersServer
}

func NewUsersServer(svcCtx *svc.ServiceContext) *UsersServer {
	return &UsersServer{
		svcCtx: svcCtx,
	}
}

func (s *UsersServer) CreateUser(ctx context.Context, in *user_grpc.UserCreateRequest) (*user_grpc.UserCreateResponse, error) {
	l := logic.NewCreateUserLogic(ctx, s.svcCtx)
	return l.CreateUser(in)
}

func (s *UsersServer) UserInfo(ctx context.Context, in *user_grpc.UserInfoRequest) (*user_grpc.UserInfoResponse, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

func (s *UsersServer) UserInfoUpdate(ctx context.Context, in *user_grpc.UserInfoUpdateRequest) (*user_grpc.UserInfoResponse, error) {
	l := logic.NewUserInfoUpdateLogic(ctx, s.svcCtx)
	return l.UserInfoUpdate(in)
}

func (s *UsersServer) UserConfUpdate(ctx context.Context, in *user_grpc.UserConfUpdateRequest) (*user_grpc.UserInfoResponse, error) {
	l := logic.NewUserConfUpdateLogic(ctx, s.svcCtx)
	return l.UserConfUpdate(in)
}

func (s *UsersServer) UserConf(ctx context.Context, in *user_grpc.UserInfoRequest) (*user_grpc.UserInfoResponse, error) {
	l := logic.NewUserConfLogic(ctx, s.svcCtx)
	return l.UserConf(in)
}

func (s *UsersServer) FriendInfo(ctx context.Context, in *user_grpc.FriendInfoRequest) (*user_grpc.FriendInfoResponse, error) {
	l := logic.NewFriendInfoLogic(ctx, s.svcCtx)
	return l.FriendInfo(in)
}

func (s *UsersServer) IsFriend(ctx context.Context, in *user_grpc.IsFriendRequest) (*user_grpc.IsFriendResponse, error) {
	l := logic.NewIsFriendLogic(ctx, s.svcCtx)
	return l.IsFriend(in)
}

func (s *UsersServer) GetFriendList(ctx context.Context, in *user_grpc.FriendListRequest) (*user_grpc.FriendListResponse, error) {
	l := logic.NewGetFriendListLogic(ctx, s.svcCtx)
	return l.GetFriendList(in)
}

// Code generated by goctl. DO NOT EDIT.
// Source: group_rpc.proto

package group

import (
	"context"

	"fim/fim_group/group_rpc/types/group_rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	MemberInfoRequest  = group_rpc.MemberInfoRequest
	MemberInfoResponse = group_rpc.MemberInfoResponse

	Group interface {
		GroupMemberInfo(ctx context.Context, in *MemberInfoRequest, opts ...grpc.CallOption) (*MemberInfoResponse, error)
	}

	defaultGroup struct {
		cli zrpc.Client
	}
)

func NewGroup(cli zrpc.Client) Group {
	return &defaultGroup{
		cli: cli,
	}
}

func (m *defaultGroup) GroupMemberInfo(ctx context.Context, in *MemberInfoRequest, opts ...grpc.CallOption) (*MemberInfoResponse, error) {
	client := group_rpc.NewGroupClient(m.cli.Conn())
	return client.GroupMemberInfo(ctx, in, opts...)
}

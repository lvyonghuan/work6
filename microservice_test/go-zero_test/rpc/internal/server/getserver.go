// Code generated by goctl. DO NOT EDIT.
// Source: get.proto

package server

import (
	"context"

	"work6/microservice_test/go-zero_test/rpc/internal/logic"
	"work6/microservice_test/go-zero_test/rpc/internal/svc"
	"work6/microservice_test/go-zero_test/rpc/rpc"
)

type GetServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedGetServer
}

func NewGetServer(svcCtx *svc.ServiceContext) *GetServer {
	return &GetServer{
		svcCtx: svcCtx,
	}
}

func (s *GetServer) Creat(ctx context.Context, in *rpc.Request) (*rpc.Response, error) {
	l := logic.NewCreatLogic(ctx, s.svcCtx)
	return l.Creat(in)
}

func (s *GetServer) GetInfo(ctx context.Context, in *rpc.Request) (*rpc.Response, error) {
	l := logic.NewGetInfoLogic(ctx, s.svcCtx)
	return l.GetInfo(in)
}

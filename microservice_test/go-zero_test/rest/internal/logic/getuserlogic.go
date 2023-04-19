package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"work6/microservice_test/go-zero_test/rest/internal/svc"
	"work6/microservice_test/go-zero_test/rest/internal/types"
	"work6/microservice_test/go-zero_test/rpc/get"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)
	good, err := l.svcCtx.GetRpc.GetInfo(l.ctx, &get.Request{
		Name: req.Name,
	})
	if err != nil {
		panic(err)
	}
	resp.Price = good.Price
	resp.Num = int(good.Num)
	return
}

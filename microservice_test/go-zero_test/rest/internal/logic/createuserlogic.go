package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"work6/microservice_test/go-zero_test/rest/internal/svc"
	"work6/microservice_test/go-zero_test/rest/internal/types"
	"work6/microservice_test/go-zero_test/rpc/get"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)
	good, err := l.svcCtx.GetRpc.Creat(l.ctx, &get.Request{
		Name:  req.Name,
		Price: req.Price,
		Num:   int32(req.Num),
	})
	resp.Num = int(good.Num)
	resp.Price = good.Price
	return resp, nil
}

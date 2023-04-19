package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"work6/microservice_test/go-zero_test/rest/internal/config"
	"work6/microservice_test/go-zero_test/rpc/get"
)

type ServiceContext struct {
	Config config.Config
	GetRpc get.Get
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		GetRpc: get.NewGet(zrpc.MustNewClient(c.GetRpc)),
	}
}

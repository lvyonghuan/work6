package logic

import (
	"context"
	"log"
	"work6/microservice_test/go-zero_test/rpc/dao"

	"work6/microservice_test/go-zero_test/rpc/internal/svc"
	"work6/microservice_test/go-zero_test/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoLogic {
	return &GetInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInfoLogic) GetInfo(in *rpc.Request) (rpy *rpc.Response, err error) {
	db := dao.NewGorm(l.svcCtx.Config.Mysql.Datasource)
	err = db.Table("store").Where("name=?", in.Name).First(&rpy).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return rpy, nil
}

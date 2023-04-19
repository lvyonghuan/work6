package logic

import (
	"context"
	"log"
	"work6/microservice_test/go-zero_test/rpc/dao"

	"work6/microservice_test/go-zero_test/rpc/internal/svc"
	"work6/microservice_test/go-zero_test/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatLogic {
	return &CreatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatLogic) Creat(in *rpc.Request) (*rpc.Response, error) {
	db := dao.NewGorm(l.svcCtx.Config.Mysql.Datasource)
	store := dao.Store{
		Name:  in.Name,
		Price: in.Price,
		Num:   int(in.Num),
	}
	err := db.Create(&store).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &rpc.Response{
		Price: store.Price,
		Num:   int32(store.Num),
	}, nil
}

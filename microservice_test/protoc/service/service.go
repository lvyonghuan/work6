package service

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"work6/microservice_test/protoc/rpc"
)

type Inter struct {
	rpc.UnimplementedServerServer
}

func (t *Inter) Serv(ctx context.Context, req *rpc.Request) (*rpc.Response, error) {
	log.Println("service:", req)
	return &rpc.Response{Result: req.Num_A + req.Num_B}, nil
}

// Service 理论上好像不该这么写，该直接写main，毕竟这个东西在别的地方，该直接在别的地方跑起来。
func Service() {
	Li, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	rpc.RegisterServerServer(server, &Inter{})
	err = server.Serve(Li)
	if err != nil {
		panic(err)
		return
	}
}

package service

import (
	"context"
	"log"
	"net"
	pb "work6/protoc/rpc"

	"google.golang.org/grpc"
)

type Inter struct {
	pb.UnimplementedServerServer
}

func (t *Inter) Serv(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Println("service:", req)
	return &pb.Response{Result: req.Num_A + req.Num_B}, nil
}

func Service() {
	Li, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterServerServer(server, &Inter{})
	err = server.Serve(Li)
	if err != nil {
		panic(err)
		return
	}
}

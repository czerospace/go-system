package main

import (
	"context"
	"go-system/stage2/grpc/grpc_test/proto"
	"google.golang.org/grpc"
	"net"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello," + request.Name,
	}, nil
}

func main() {
	// 1.实例化 grpc server
	g := grpc.NewServer()
	// 2.注册
	proto.RegisterGreeterServer(g, &Server{})
	// 3.启动服务
	lis, err := net.Listen("tcp", "0.0.0.0:8088")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc: " + err.Error())
	}
}

package main

import (
	"go-system/stage2/grpc/stream_grpc_test/proto"
	"google.golang.org/grpc"
	"net"
)

const PORT = ":50052"

type server struct {
}

func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	return nil
}
func (s *server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	return nil
}
func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
}

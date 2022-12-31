package main

import (
	"context"
	"fmt"
	"go-system/stage2/grpc/grpc_test/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "winnie"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}

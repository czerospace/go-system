package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"

	"go-system/stage2/grpc/grpc_interpretor/proto"
)

func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("method=%s req=%v rep=%v duration=%s error=%v\n", method, req, reply, time.Since(start), err)
	return err
}

func main() {
	//stream
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	// 指定客户端interceptor
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))

	conn, err := grpc.Dial("localhost:50051", opts...)
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

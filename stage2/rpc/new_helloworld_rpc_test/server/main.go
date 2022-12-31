package main

import (
	"go-system/stage2/rpc_test/new_helloworld_rpc_test/handler"
	"go-system/stage2/rpc_test/new_helloworld_rpc_test/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	// 1.实例化一个 server
	listener, _ := net.Listen("tcp", ":1234")
	// 2.注册处理逻辑 handler
	_ = server_proxy.RegisterHelloService(&handler.NewHelloService{})
	// 3.启动服务
	for {
		conn, _ := listener.Accept() // 当一个新的链接进来的时候
		rpc.ServeConn(conn)
	}

}

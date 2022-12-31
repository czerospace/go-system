package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	// 返回值事通过修改 reply 的值
	*reply = "Hello," + request
	return nil
}

func main() {
	// 1.实例化一个 server
	listener, _ := net.Listen("tcp", ":1234")
	// 2.注册处理逻辑 handler
	_ = rpc.RegisterName("HelloService", &HelloService{})
	// 3.启动服务
	conn, _ := listener.Accept() // 当一个新的链接进来的时候
	rpc.ServeConn(conn)
}

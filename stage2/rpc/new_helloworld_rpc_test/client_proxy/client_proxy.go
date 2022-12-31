package client_proxy

import (
	"go-system/stage2/rpc_test/new_helloworld_rpc_test/handler"
	"net/rpc"
)

type HelloServerStub struct {
	*rpc.Client
}

// 在 go 语言中没有类、对象 就意味着没有初始化方法，通常用 New+xxx 定义函数来初始化
func NewHelloServiceClient(protol, address string) HelloServerStub {
	conn, err := rpc.Dial(protol, address)
	if err != nil {
		panic("connect error!")
	}
	return HelloServerStub{conn}
}

// 封装 Hello 方法
func (c *HelloServerStub) Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}

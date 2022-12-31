package main

import (
	"fmt"
	"go-system/stage2/rpc_test/new_helloworld_rpc_test/client_proxy"
)

func main() {
	// 1.建立链接
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")

	var reply string // string 有默认值
	err := client.Hello("winnie", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}

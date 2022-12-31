package main

import (
	"fmt"
	helloworld "go-system/stage2/rpc_test/helloworld_rpc_test/proto"
	"google.golang.org/protobuf/proto"
)

func main() {
	req := helloworld.HelloRequest{
		Name:    "winnie",
		Age:     18,
		Courses: []string{"go", "gin", "微服务"},
	}
	rsp, _ := proto.Marshal(&req)
	newReq := helloworld.HelloRequest{}
	_ = proto.Unmarshal(rsp, &newReq)
	fmt.Println(newReq.Name, newReq.Age, newReq.Courses)
}

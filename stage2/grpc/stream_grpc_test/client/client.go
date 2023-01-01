package main

import (
	"context"
	"fmt"
	"go-system/stage2/grpc/stream_grpc_test/proto"
	"google.golang.org/grpc"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 服务端流模式
	c := proto.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "winnie"})
	for {
		a, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a.Data)
	}

	// 客户端流模式
	putS, _ := c.PutStream(context.Background())
	i := 0
	for {
		i++
		_ = putS.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("winnie+%d", i),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	// 双向流模式
	// 获取双向流通道
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{} // 实例化协程
	wg.Add(2)              // 开启2个协程
	// 接收
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息: " + data.Data)
		}
	}()
	// 发送
	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamReqData{Data: "我是客户端"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}

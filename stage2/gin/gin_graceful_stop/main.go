package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	/*
			优雅退出，当我们关闭程序的时候应该做的后续处理
			比如微服务 启动之前或者启动之后会 将当前服务的ip地址和端口号注册到注册中心
		    优雅退出 要在服务停止的时候 告知注册中心，从注册中心下线
	*/
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	// 使用协程监听端口
	go func() {
		err := router.Run(":8083")
		if err != nil {
			panic(err)
		}
	}()

	// 定义一个 chan 接收 系统信号 无法接收 kill -9
	quit := make(chan os.Signal)
	// 处理 ctrl c 和 kill 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 读取管道内容，只要有数据进来，就代表要退出
	<-quit
	// 处理后续的逻辑,用打印模拟
	fmt.Println("关闭server中...")
	fmt.Println("注销服务中...")
}

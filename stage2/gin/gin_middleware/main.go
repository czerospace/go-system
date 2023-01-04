package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// MyLogger 自定义一个中间件，打印耗时和状态
func MyLogger() gin.HandlerFunc {
	// 固定用法,返回一个 gin.HandlerFunc 函数
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "123456")
		// 让原本该执行的逻辑继续执行
		c.Next()
		end := time.Since(t)
		fmt.Printf("耗时:%v\n", end)
		status := c.Writer.Status()
		fmt.Println("状态:", status)
	}
}

/*
	需求：希望 在验证登录之后才能返回 "message": "pong"
*/

// TokenRequired 验证 header 中的 token
func TokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		// 获取 header
		for k, v := range c.Request.Header {
			// 如果 header 是 x-token 就取 header 的值 赋给 token
			if k == "X-Token" {
				token = v[0]
			}
			fmt.Println(k, v)
		}
		if token != "winnie" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未登录",
			})
			// 这里必须要用 c.Abort() ，用 return 无效
			// gin 本身有个 index，return 之后 index会++，进入下一个中间件
			c.Abort()
		}
		c.Next()
	}
}

func main() {
	router := gin.Default()
	// 使用自定义的中间件 MyLogger()
	router.Use(MyLogger(), TokenRequired())

	router.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err := router.Run(":8083")
	if err != nil {
		panic(err)
	}
}

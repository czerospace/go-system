package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	router := gin.Default()

	// 指定模板的目录
	// LoadHTMLFiles 会加载指定的文件
	// 相对路径在 goland 中无法加载到，打印下当前路径看下
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0])) // C:\Users\czero\AppData\Local\Temp\GoLand
	fmt.Println(dir)

	// 加载静态文件
	// <link rel="stylesheet" href="/static/style.css">
	router.Static("/static", "./static")

	// 在终端中编译成为二进制文件后使用相对路径
	// 加载单个文件
	//router.LoadHTMLFiles("templates/index.tmpl", "templates/goods.html")
	// 加载所有文件
	router.LoadHTMLGlob("templates/**/*")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "winniethepooh",
		})
	})

	router.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"name": "winnie",
		})
	})

	/*
		goods/list.html 和 users/list.html
		文件同名，需要在模板中使用 define 定义处理
	*/
	router.GET("/goods/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods/list.html", gin.H{
			"name": "winnie",
		})
	})
	router.GET("/users/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/list.html", gin.H{
			"name": "winnie",
		})
	})

	err := router.Run(":8083")
	if err != nil {
		return
	}
}

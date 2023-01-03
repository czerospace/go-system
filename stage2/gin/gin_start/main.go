package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// gin.H 实际上是一个 map[string]string
func pingpong(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "pingpong",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", pong)
	r.GET("/pingpong", pingpong)
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		return
	}
}

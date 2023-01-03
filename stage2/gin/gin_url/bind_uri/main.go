package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Person 使用结构体约束参数
type Person struct {
	ID   int    `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.Status(404)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name": person.Name,
			"id":   person.ID,
		})
	})
	err := router.Run(":8083")
	if err != nil {
		panic(err)
	}
}

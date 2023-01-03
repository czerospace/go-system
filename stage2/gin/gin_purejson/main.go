package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// Serves unicode entities
	router.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// Serves literal characters
	router.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// listen and serve on 0.0.0.0:8080
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-system/stage2/gin/gin_json_protobuf/proto"
)

func main() {
	router := gin.Default()
	router.GET("/moreJSON", moreJSON)
	router.GET("/someProtoBuf", returnProto)
	err := router.Run(":8083")
	if err != nil {
		panic(err)
	}

}

func moreJSON(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "winnie"
	msg.Message = "这是一个测试json"
	msg.Number = 9527
	c.JSON(http.StatusOK, msg)
}

func returnProto(c *gin.Context) {
	course := []string{"go", "python", "微服务"}
	user := &protogin.Teacher{
		Name:   "winnie",
		Course: course,
	}
	c.ProtoBuf(http.StatusOK, user)
}

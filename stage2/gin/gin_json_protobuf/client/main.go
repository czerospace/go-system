package main

import (
	"fmt"
	protogin "go-system/stage2/gin/gin_json_protobuf/proto"
	"google.golang.org/protobuf/proto"

	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8083/someProtoBuf")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var teacher protogin.Teacher
	err = proto.Unmarshal(body, &teacher)
	fmt.Println(teacher.Name, teacher.Course)

}

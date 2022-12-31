package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// http://127.0.0.1:8000/add?a=1&b=2
	// 格式化返回: json {"data":3}
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm() // 解析参数
		fmt.Println("path: ", r.URL.Path)
		// 从 URL 获取取 a b 并转成 int 类型
		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])
		// 设置 Header 信息
		w.Header().Set("Content-Type", "application/json")
		// 序列化为 json 格式
		JData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		_, _ = w.Write(JData)
	})

	_ = http.ListenAndServe(":8000", nil)
}

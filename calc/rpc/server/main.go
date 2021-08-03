package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	// http://127.0.0.1:8000/add?a=1&b=2
	// 返回的格式化：json {"data":3}
	/*
		1. callID 的问题：r.URL.Path
		2. 数据的传输协议： url的参数传递协议
		3. 网络传输协议 http
	*/
	http.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm() // 解析参数
		fmt.Println("path: ", request.URL.Path)
		a, _ := strconv.Atoi(request.Form["a"][0]) // string 类型转 int 类型
		b, _ := strconv.Atoi(request.Form["b"][0]) // string 类型转 int 类型

		writer.Header().Set("Content-Type", "application/json")
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		_, _ = writer.Write(jData)
	})

	http.ListenAndServe(":8000", nil)
}

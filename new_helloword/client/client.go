package main

import (
	"fmt"
	"learngoframework/new_helloword/client_proxy"
)

func main() {

	// 1. 只想写业务逻辑，不想关注每个函数的名称

	// 1. 建立链接
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")

	// 2. 发起调用
	//var reply *string = new(string)
	var reply string // string 有默认值
	err := client.Hello("bobby", &reply)
	if err != nil {
		fmt.Println("调用失败")
		panic(err)
	}

	// 3. 输出调用结果
	//fmt.Println(*reply)
	fmt.Println(reply)

	/*
		1. 这些概念在grpc中都有对应
		2. 发自灵魂的拷问？ server_proxy 和 client_proxy 能否自动生成 为多种语言生成
		以上问题都可以满足 这个就是 protobuf + grpc
	*/

}

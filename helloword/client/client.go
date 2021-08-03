package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1. 建立链接
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("连接失败")
		panic(err)
	}

	// 2. 发起调用
	//var reply *string = new(string)
	var reply string // string 有默认值
	err = client.Call("HelloService.Hello", "bobby", &reply)
	if err != nil {
		fmt.Println("调用失败")
		panic(err)
	}

	// 3. 输出调用结果
	//fmt.Println(*reply)
	fmt.Println(reply)

}

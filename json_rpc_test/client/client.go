package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 1. 建立链接
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("连接失败")
		panic(err)
	}

	// 2. 发起调用
	//var reply *string = new(string)
	var reply string                                               // string 有默认值
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn)) //设置客户端序列化协议为 json
	err = client.Call("HelloService.Hello", "bobby", &reply)
	if err != nil {
		fmt.Println("调用失败")
		panic(err)
	}

	// 3. 输出调用结果
	//fmt.Println(*reply)
	fmt.Println(reply)

}

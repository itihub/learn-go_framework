package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learngoframework/grpc_test/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure()) // 拨号 建立连接
	if err != nil {
		panic(err)
	}
	defer conn.Close() // 关闭连接

	c := proto.NewGreeterClient(conn) // 生成客户端

	r, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "bobby",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Message)
}

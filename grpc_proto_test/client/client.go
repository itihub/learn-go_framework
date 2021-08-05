package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learngoframework/grpc_proto_test/client/proto-bak"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50053", grpc.WithInsecure()) // 拨号 建立连接
	if err != nil {
		panic(err)
	}
	defer conn.Close() // 关闭连接

	c := proto_bak.NewGreeterClient(conn) // 生成客户端

	r, err := c.SayHello(context.Background(), &proto_bak.HelloRequest{
		Name: "bobby",
		Url:  "http://example.com",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Message)
}

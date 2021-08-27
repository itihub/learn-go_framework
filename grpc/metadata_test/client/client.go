package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	proto2 "learngoframework/grpc/grpc_test/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure()) // 拨号 建立连接
	if err != nil {
		panic(err)
	}
	defer conn.Close() // 关闭连接

	c := proto2.NewGreeterClient(conn) // 生成客户端

	// 生成metadata方式一
	//md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	// 生成metadata方式二
	md := metadata.New(map[string]string{
		"name":     "bobby",
		"password": "123456",
	})
	// 将metadata加入到Context中
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	r, err := c.SayHello(ctx, &proto2.HelloRequest{
		Name: "bobby",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Message)
}

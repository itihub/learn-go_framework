package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	proto2 "learngoframework/grpc/grpc_proto_test/client/proto"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50053", grpc.WithInsecure()) // 拨号 建立连接
	if err != nil {
		panic(err)
	}
	defer conn.Close() // 关闭连接

	c := proto2.NewGreeterClient(conn) // 生成客户端

	r, err := c.SayHello(context.Background(), &proto2.HelloRequest{
		Name: "bobby",
		Url:  "http://example.com",
		G:    proto2.Gender_FEMALE, // 使用枚举类型
		Mp: map[string]string{
			"key": "value",
		},
		AddTime: timestamppb.New(time.Now()), // protobuf 内建时间戳类型
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Message)
}

package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	proto2 "learngoframework/grpc/grpc_test/proto"
	"net"
)

type Server struct {
}

// 业务逻辑 提供接口/函数
func (s *Server) SayHello(ctx context.Context, request *proto2.HelloRequest) (*proto2.HelloReply, error) {

	md, ok := metadata.FromIncomingContext(ctx) // 从context中获取metadata
	if ok {
		fmt.Println("get metadata error")
	}

	for key, value := range md {
		fmt.Println(key, value)
	}

	if nameSlice, ok := md["name"]; ok {
		fmt.Println(nameSlice)
		for i, e := range nameSlice {
			fmt.Println(i, e)
		}
	}

	return &proto2.HelloReply{
		Message: "hello, " + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()                      // 实例化 GRPC 服务
	proto2.RegisterGreeterServer(g, &Server{}) // 注册服务接口

	lis, err := net.Listen("tcp", "0.0.0.0:50051") // 监听端口
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	err = g.Serve(lis) // 启动服务
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}

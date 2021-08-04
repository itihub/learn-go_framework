package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learngoframework/grpc_proto_test/server/proto"
	"net"
)

type Server struct {
}

// 业务逻辑 提供接口/函数
func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	fmt.Println(request.Name)
	fmt.Println(request.Url)
	// url 与 name 值反了 是因为客户端和服务器端的proto文件的参数顺序不一致导致了
	return &proto.HelloReply{
		Message: "hello, " + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()                     // 实例化 GRPC 服务
	proto.RegisterGreeterServer(g, &Server{}) // 注册服务接口

	lis, err := net.Listen("tcp", "0.0.0.0:50053") // 监听端口
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	err = g.Serve(lis) // 启动服务
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}

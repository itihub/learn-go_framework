package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	proto2 "learngoframework/grpc/grpc_proto_test/server/proto"
	"net"
)

type Server struct {
}

// 业务逻辑 提供接口/函数
func (s *Server) SayHello(ctx context.Context, request *proto2.HelloRequest) (*proto2.HelloReply, error) {
	fmt.Println(request.Name)
	fmt.Println(request.Url)
	// url 与 name 值反了 是因为客户端和服务器端的proto文件的参数顺序不一致导致了
	return &proto2.HelloReply{
		Message: "hello, " + request.Name,
		Data: []*proto2.HelloReply_Result{
			&proto2.HelloReply_Result{
				Name: request.Name,
				Url:  "",
			},
		},
	}, nil
}

func (s *Server) Ping(ctx context.Context, in *emptypb.Empty) (*proto2.Pong, error) {
	return &proto2.Pong{
		Id: "1",
	}, nil
}

func main() {
	g := grpc.NewServer()                      // 实例化 GRPC 服务
	proto2.RegisterGreeterServer(g, &Server{}) // 注册服务接口

	lis, err := net.Listen("tcp", "0.0.0.0:50053") // 监听端口
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	err = g.Serve(lis) // 启动服务
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}

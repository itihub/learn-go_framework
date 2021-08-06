package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learngoframework/grpc_error_test/proto"
	"net"
)

type Server struct {
}

// 业务逻辑 提供接口/函数
func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {

	//return nil, status.Error(codes.InvalidArgument, "invalid username")
	return nil, status.Errorf(codes.NotFound, "记录未找到:%s", request.Name)

	//return &proto.HelloReply{
	//	Message: "hello, " + request.Name,
	//}, nil
}

func main() {
	g := grpc.NewServer()                     // 实例化 GRPC 服务
	proto.RegisterGreeterServer(g, &Server{}) // 注册服务接口

	lis, err := net.Listen("tcp", "0.0.0.0:8080") // 监听端口
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	err = g.Serve(lis) // 启动服务
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}

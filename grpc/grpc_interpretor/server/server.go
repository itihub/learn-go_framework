package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	proto2 "learngoframework/grpc/grpc_interpretor/proto"
	"net"
)

type Server struct {
}

// 业务逻辑 提供接口/函数
func (s *Server) SayHello(ctx context.Context, request *proto2.HelloRequest) (*proto2.HelloReply, error) {
	return &proto2.HelloReply{
		Message: "hello, " + request.Name,
	}, nil
}

func main() {

	// 自定义Server端拦截器
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("接收到一个新的请求")
		res, err := handler(ctx, req)
		fmt.Println("请求完成")
		return res, err
	}

	otp := grpc.UnaryInterceptor(interceptor)  //一元调用grpc拦截器 注册
	g := grpc.NewServer(otp)                   // 实例化 GRPC 服务
	proto2.RegisterGreeterServer(g, &Server{}) // 注册服务接口

	lis, err := net.Listen("tcp", "0.0.0.0:8080") // 监听端口
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	err = g.Serve(lis) // 启动服务
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}

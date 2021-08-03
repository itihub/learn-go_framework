package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	// 返回值是通过修改reply的值
	*reply = "hello, " + request
	return nil
}

func main() {
	// 1. 实例化一个server
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	// 2. 注册处理逻辑 handler 将服务注册RPC中
	_ = rpc.RegisterName("HelloService", &HelloService{})

	// 3. 启动服务
	conn, err := listen.Accept() // 当一个新的链接进来的时候，套接字传入rpc处理
	if err != nil {
		panic(err)
	}
	rpc.ServeConn(conn)
}

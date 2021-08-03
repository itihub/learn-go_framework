package main

import (
	"learngoframework/new_helloword/handler"
	"learngoframework/new_helloword/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	// 1. 实例化一个server
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	// 2. 注册处理逻辑 handler 将服务注册RPC中
	server_proxy.RegisterHelloService(&handler.HelloService{})

	// 3. 启动服务
	for {
		conn, err := listen.Accept() // 当一个新的链接进来的时候，套接字传入rpc处理
		if err != nil {
			panic(err)
		}
		go rpc.ServeConn(conn)
	}

}

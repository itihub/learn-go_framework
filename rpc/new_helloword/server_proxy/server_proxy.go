package server_proxy

import (
	handler2 "learngoframework/rpc/new_helloword/handler"
	"net/rpc"
)

type HelloServicer interface {
	Hello(request string, reply *string) error
}

// 如果做到解耦 - 我们关系的是函数 鸭子类型
func RegisterHelloService(srv HelloServicer) error {
	return rpc.RegisterName(handler2.HelloServiceName, srv)
}

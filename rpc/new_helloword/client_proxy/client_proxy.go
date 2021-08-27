package client_proxy

import (
	"fmt"
	handler2 "learngoframework/rpc/new_helloword/handler"
	"net/rpc"
)

/*
	代理类
*/

type HelloServiceStub struct {
	*rpc.Client
}

// 在go语言中 没有类、对象 就意味着没有初始化方法
func NewHelloServiceClient(protol, address string) HelloServiceStub {
	conn, err := rpc.Dial(protol, address)
	if err != nil {
		fmt.Println("connect error!")
		panic(err)
	}
	return HelloServiceStub{conn}
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handler2.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}

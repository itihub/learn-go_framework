package main

import (
	"fmt"
	"google.golang.org/grpc"
	"learngoframework/stream_grpc_test/proto"
	"net"
	"sync"
	"time"
)

const PORT = ":50002"

type server struct {
}

// 业务逻辑 提供接口/函数
func (s *server) GetStream(request *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

func (s *server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	for {
		a, err := cliStr.Recv()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(a.Data)
	}
	return nil
}

func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息：" + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamResData{
				Data: "我是服务器",
			})

			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", PORT) // 监听端口
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	g := grpc.NewServer()                     // 实例化 GRPC 服务
	proto.RegisterGreeterServer(g, &server{}) // 注册服务接口
	err = g.Serve(lis)                        // 启动服务
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}

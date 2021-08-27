package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	proto2 "learngoframework/grpc/stream_grpc_test/proto"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50002", grpc.WithInsecure()) // 拨号 建立连接
	if err != nil {
		panic(err)
	}
	defer conn.Close() // 关闭连接

	c := proto2.NewGreeterClient(conn) // 生成客户端

	// 服务端流模式
	res, err := c.GetStream(context.Background(), &proto2.StreamReqData{Data: "Server-side"})
	if err != nil {
		panic(err)
	}
	for {
		a, err := res.Recv() // 接收数据
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a.Data)
	}

	// 客户端流模式
	putS, _ := c.PutStream(context.Background())
	i := 0
	for {
		i++
		putS.Send(&proto2.StreamReqData{
			Data: fmt.Sprintf("Client-side %d", i),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	// 双向流模式
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到服务器消息：" + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto2.StreamReqData{
				Data: "我是客户端",
			})

			time.Sleep(time.Second)
		}
	}()

	wg.Wait()

}

package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	proto2 "learngoframework/grpc/grpc_token_auth_test/proto"
	"time"
)

func main() {

	// 自定义客户端拦截器
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		// 为每次请求统一添加metadata信息
		md := metadata.New(map[string]string{
			"app_id":  "101010",
			"app_key": "123456",
		})
		ctx = metadata.NewOutgoingContext(context.Background(), md)
		err := invoker(ctx, method, req, reply, cc, opts...) // 继续向下调用
		fmt.Printf("耗时：%s\n", time.Since(start))
		return err
	}

	// 注册拦截器写法方式一
	otp := grpc.WithUnaryInterceptor(interceptor)                      // 注册拦截器
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), otp) // 拨号 建立连接

	// 注册拦截器写法方式二
	//var otps []grpc.DialOption
	//otps = append(otps, grpc.WithInsecure())
	//otps = append(otps, grpc.WithUnaryInterceptor(interceptor))
	//conn, err := grpc.Dial("127.0.0.1:8080", otps...)

	if err != nil {
		panic(err)
	}
	defer conn.Close() // 关闭连接

	c := proto2.NewGreeterClient(conn) // 生成客户端

	r, err := c.SayHello(context.Background(), &proto2.HelloRequest{
		Name: "bobby",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Message)
}

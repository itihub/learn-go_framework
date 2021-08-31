package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"learngoframework/grpc/grpc_interpretor/proto"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
)

func main() {

	// 自定义客户端拦截器
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...) // 继续向下调用
		fmt.Printf("耗时：%s\n", time.Since(start))
		return err
	}

	// 全局设置重试
	retryOpts := []grpc_retry.CallOption{
		grpc_retry.WithMax(3),                                                          // 重试次数
		grpc_retry.WithPerRetryTimeout(1 * time.Second),                                // 超时时间
		grpc_retry.WithCodes(codes.Unknown, codes.DeadlineExceeded, codes.Unavailable), // 重试条件
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))                                     // 注册拦截器
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOpts...))) // 注册重试拦截器 问题？设置超时时间 重试次数 以及 服务器返回什么状态码进行重试(判断重试条件)
	conn, err := grpc.Dial("127.0.0.1:8080", opts...)                                               // 拨号 建立连接

	if err != nil {
		panic(err)
	}
	defer conn.Close() // 关闭连接

	c := proto.NewGreeterClient(conn) // 生成客户端

	r, err := c.SayHello(
		context.Background(),
		&proto.HelloRequest{Name: "bobby"},

	// 局部设置重试
	//grpc_retry.WithMax(3), // 重试次数
	//grpc_retry.WithPerRetryTimeout(1 * time.Second), // 超时时间
	//grpc_retry.WithCodes(codes.Unknown, codes.DeadlineExceeded, codes.Unavailable), // 重试条件
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Message)
}

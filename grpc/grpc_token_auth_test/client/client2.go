package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	proto2 "learngoframework/grpc/grpc_token_auth_test/proto"
)

type customCredential struct {
}

// 实现credentials.PerRPCCredentials接口
func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"app_id":  "101010",
		"app_key": "123456",
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {

	// 设置grpc认证信息
	otp := grpc.WithPerRPCCredentials(customCredential{})

	// 注册拦截器写法方式一
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

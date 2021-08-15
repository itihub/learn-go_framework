package main

import (
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // Is’s important
	"google.golang.org/grpc"
	"learngoframework/grpc_lb_test/proto"
	"log"
)

// grpc 负载均衡
func main() {
	conn, err := grpc.Dial(
		"consul://local.docker.node1.com:8500/user_srv?wait=14s&tag=",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewUserClient(conn)

	for i := 0; i < 10; i++ {
		rsp, err := client.GetUserList(context.Background(), &proto.PageInfo{
			Pn:    1,
			PSize: 5,
		})
		if err != nil {
			panic(err)
		}

		for index, data := range rsp.Data {
			fmt.Println(index, data)
		}
	}

}

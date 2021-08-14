package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

// consul注册服务
func Register(address string, port int, name string, tags []string, id string) error {
	// 设置配置
	cfg := api.DefaultConfig()
	cfg.Address = "local.docker.node1.com:8500"

	// 生成consul客户端
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address

	// 生成健康检查对象 HTTP方式
	check := &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%v:%d/health", address, port), // 检查地址
		Timeout:                        "5s",                                              // 超时时间
		Interval:                       "5s",                                              // 检查间隔
		DeregisterCriticalServiceAfter: "10s",
	}
	registration.Check = check

	// 进行注册
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

// 获取consul所有注册服务
func AllServices() {
	// 设置配置
	cfg := api.DefaultConfig()
	cfg.Address = "local.docker.node1.com:8500"

	// 生成consul客户端
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)
	}
}

// 指定Name获取服务列表
func FilterService(name string) {
	// 设置配置
	cfg := api.DefaultConfig()
	cfg.Address = "local.docker.node1.com:8500"

	// 生成consul客户端
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, name))
	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)
	}
}

func main() {

	err := Register("192.168.0.10", 8021, "user-web", []string{"shop", "user-web"}, "user-web")
	if err != nil {
		panic(err)
	}

	AllServices()
	FilterService("user-web")
}

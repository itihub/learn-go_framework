package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"time"
)

func main() {

	// 创建serverConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "local.docker.node1.com",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}

	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "user",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
	}

	// 从nacos读取配置文件
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user-web.yaml",
		Group:  "dev",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(content) // yaml 格式字符串

	//监听配置变化
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "user-web.yaml",
		Group:  "dev",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("配置文件产生变化...")
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	if err != nil {
		panic(err)
	}

	time.Sleep(3000 * time.Second)

}

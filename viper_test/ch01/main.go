package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	ServerName string `mapstructure:"name"` // 配置文件中配置项映射
	Port       int    `mapstructure:"port"` // 配置文件中配置项映射
}

func main() {
	v := viper.New()
	// 设置读取文件名称
	//v.SetConfigFile("config.yaml")
	v.SetConfigFile("viper_test/ch01/config.yaml")
	// 读取
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	// 获取配置文件变量
	fmt.Printf("%v\n", v.Get("name"))

	// 配置文件映射struct
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)

}

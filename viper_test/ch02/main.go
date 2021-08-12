package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"time"

	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// 嵌套配置文件
type ServerConfig struct {
	ServerName string      `mapstructure:"name"` // 配置文件中配置项映射
	Port       int         `mapstructure:"port"` // 配置文件中配置项映射
	MysqlInfo  MysqlConfig `mapstructure:"mysql"`
}

// 读取环境变量 系统or go env 参数为变量名称
func getEnvInfo(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}

func main() {

	// 读取环境变量值
	fmt.Println(getEnvInfo("SHOP_DEBUG"))
	env := getEnvInfo("SHOP_DEBUG")

	configFilePrefix := "config"
	configFileName := fmt.Sprintf("viper_test/ch02/%s-debug.yaml", configFilePrefix)
	if env == "debug" {
		configFileName = fmt.Sprintf("viper_test/ch02/%s-debug.yaml", configFilePrefix)
	}
	v := viper.New()
	// 设置读取文件名称
	//v.SetConfigFile("config-debug.yaml")
	v.SetConfigFile(configFileName)
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

	//	viper功能 - 动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file channed: ", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&serverConfig)
		fmt.Println(serverConfig)
	})

	time.Sleep(time.Second * 300)
}

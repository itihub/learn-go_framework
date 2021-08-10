package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {

	// gin 优雅退出:当我们关闭程序应该如何做后续处理
	// 微服务 启动之前或者启动之后会做一件事:将当前的服务的IP地址和端口号注册到注册中心
	// 我们当前的服务停止了以后并没有告知注册中心
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	// 协程启动
	go func() {
		router.Run(":8083")
	}()

	// 如果想要接收关闭信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 监听到ctrl+c信号或者kill信号 发送消息到chan 注意：kill -9 强杀是监听不到命令的
	<-quit                                               // 接收消息 无则阻塞 有责向下执行

	// 处理后续逻辑
	fmt.Println("关闭server中...")
	fmt.Println("注销服务...")
}

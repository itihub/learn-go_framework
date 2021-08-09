package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func pong(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	// 实例化Gin的server对象
	r := gin.Default()
	r.GET("/ping", pong)

	r.Run(":8083") // 运行 默认监听8080端口
}

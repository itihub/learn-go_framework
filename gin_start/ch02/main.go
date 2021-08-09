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
	// 使用默认方式创建gin的路由器 会开启logger 和 recovery（crash-free） 中间件
	router := gin.Default()
	// 使用New方式创建gin的路由器 不会开启logger 和 recovery（crash-free） 中间件 也就是无日志和异常恢复功能
	//router := gin.New()

	// restful 开发
	router.GET("/pingGet", pong)
	router.POST("/pingPost", pong)
	router.PUT("/pingPut", pong)
	router.DELETE("/pingDelete", pong)
	router.PATCH("/pingPatch", pong)
	router.HEAD("/pingHead", pong)
	router.OPTIONS("/pingOptions", pong)

	router.Run(":8083") // 运行 默认监听8080端口
}

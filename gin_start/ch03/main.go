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

	// 未分组 存在大量的path重复
	//router.GET("/goods/list", goodsList)
	//router.GET("/goods/1", goodsDetail)
	//router.POST("/goods/add", createGoods)

	// 使用分组方式
	goodsGroup := router.Group("/goods")
	{
		goodsGroup.GET("/list", goodsList)
		goodsGroup.GET("/1", goodsDetail)
		goodsGroup.POST("/add", createGoods)
	}

	// 路由分组
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8083") // 运行 默认监听8080端口
}

func readEndpoint(context *gin.Context) {

}

func submitEndpoint(context *gin.Context) {

}

func loginEndpoint(context *gin.Context) {

}

func createGoods(context *gin.Context) {

}

func goodsDetail(context *gin.Context) {

}

func goodsList(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"name": "goodsList",
	})
}

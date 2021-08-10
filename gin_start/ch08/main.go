package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义中间件
func MyLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()

		context.Set("example", "123456")
		context.Next() // 让原本该执行的逻辑继续向下执行

		end := time.Since(t) // 计算时间差
		fmt.Printf("耗时:%v\n", end)

		status := context.Writer.Status()
		fmt.Println("状态", status)
	}
}

// 自定义中间件 token验证
func AuthRequired() gin.HandlerFunc {
	return func(context *gin.Context) {
		var token string
		for k, v := range context.Request.Header {
			if k == "X-Token" {
				token = v[0]
			}
			fmt.Println(k, v)
		}

		if token != "abcd" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未登录",
			})
			// 终止中间件后续逻辑 使用return无效 需要使用context.Abort()
			context.Abort()
		}
		context.Next()
	}
}

func main() {
	router := gin.New()

	// 全局使用中间件
	// 使用logger中间件
	router.Use(gin.Logger())
	// 使用Recovery间件
	router.Use(gin.Recovery())

	router.Use(AuthRequired())

	// 方式二
	//router.Use(gin.Logger(), gin.Recovery())

	// 为一组 添加中间件
	authorized := router.Group("/goods")
	authorized.Use(MyLogger())
	{
		authorized.GET("/ping", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	router.Run(":8083")
}

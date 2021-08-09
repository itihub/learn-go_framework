package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	//ID string `uri:"id" binding:"required,uuid"`	// 参数约束：参数来源URL，必传 格式为UUID
	ID   int    `uri:"id" binding:"required"`   // 参数约束：参数来源URL，必传
	Name string `uri:"name" binding:"required"` //	参数约束：参数来源URL，必传
}

func main() {
	router := gin.Default()

	// 定义RL中的变量
	//router.GET("/goods/:id/:action", goodsDetail)
	router.GET("/goods/:id/*action", goodsDetail) // *号可以获取action以后的路径地址

	router.GET("/:name/:id", func(context *gin.Context) {
		var person Person
		// 参数校验以及绑定
		if err := context.ShouldBindUri(&person); err != nil {
			context.Status(http.StatusNotFound)
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"name": person.Name,
			"id":   person.ID,
		})
	})

	router.Run(":8083")
}

func goodsDetail(context *gin.Context) {
	// 获取url变量 无法限制参数类型
	id := context.Param("id")
	action := context.Param("action")
	context.JSON(http.StatusOK, gin.H{
		"id":     id,
		"action": action,
	})
}

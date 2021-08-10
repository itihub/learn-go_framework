package main

import (
	"github.com/gin-gonic/gin"
	"learngoframework/gin_start/ch06/proto"
	"net/http"
)

func main() {

	router := gin.Default()

	// JSON 输出格式
	router.GET("/moreJSON", moreJSON)
	// Protobuf 输出格式
	router.GET("/someProtoBuf", returnProto)
	// JSON返回会将特殊的HTML字符替换为对应的unicode字符，比如 < 替换为 \u003c,如果想原样输出html,使用PureJSON
	router.GET("/purejson", pureJSON)
	router.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"html": "<b>Hello, World</b>",
		})
	})

	router.Run(":8083")
}

func pureJSON(context *gin.Context) {
	context.PureJSON(http.StatusOK, gin.H{
		"html": "<b>Hello, World</b>",
	})
}

func returnProto(context *gin.Context) {
	course := []string{"python", "go", "微服务"}
	context.ProtoBuf(http.StatusOK, &proto.Teacher{
		Name:   "jimmy",
		Course: course,
	})
}

func moreJSON(context *gin.Context) {
	var msg struct {
		Name    string `json:"name"`
		Message string
		Number  int
	}

	msg.Name = "jimmy"
	msg.Message = "测试JSON"
	msg.Number = 20

	context.JSON(http.StatusOK, msg)
}

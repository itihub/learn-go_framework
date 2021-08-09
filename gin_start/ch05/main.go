package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	// 匹配url格式：/welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", welcome)

	router.POST("/form_post", formPost)

	router.POST("/post", getPost)

	router.Run(":8083")
}

func getPost(context *gin.Context) {
	id := context.Query("id")
	page := context.DefaultQuery("page", "0")
	name := context.PostForm("name")
	message := context.DefaultPostForm("message", "信息")
	context.JSON(http.StatusOK, gin.H{
		"id":      id,
		"page":    page,
		"name":    name,
		"message": message,
	})
}

func formPost(context *gin.Context) {
	message := context.PostForm("message")               // 从POST表单取值
	nick := context.DefaultPostForm("nick", "anonymous") // 从POST表单取值并设置未取到的默认值
	context.JSON(http.StatusOK, gin.H{
		"message": message,
		"nick":    nick,
	})
}

func welcome(context *gin.Context) {
	firstname := context.DefaultQuery("firstname", "jimmy") // DefaultQuery 取值必须设置默认值
	lastname := context.Query("lastname")                   // Query方式取值，无需设置默认值
	// 字符串格式返回
	context.String(http.StatusOK, "firstname：%s lastname:%s", firstname, lastname)
}

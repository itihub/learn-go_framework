package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// 获取当前路径 goland运行会在{{sys.User}}\AppData\Local\Temp 临时路径所以相对路径会生效
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(dir)

	router.Static("/static", "./static") // 静态文件处理 第一个参数：请求路径;第二个参数：相对存放路径

	// 加载目录下的所有文件
	//router.LoadHTMLGlob("templates/*")

	// 加载目录下的所有目录的文件
	router.LoadHTMLGlob("templates/**/*")

	// LoadHTMLFiles会将指定的目录下的文件加载好 相对路径
	// 为什么通过goland运行main.go没有生成main.exe文件
	//router.LoadHTMLFiles("templates/index.tmpl", "templates/goods.html")

	// 返回HTML
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "模板测试",
		})
	})

	router.GET("/goods", func(context *gin.Context) {
		context.HTML(http.StatusOK, "goods.html", gin.H{
			"name": "微服务开发",
		})
	})

	// 如果没有在模板中使用define定义 那么使用默认的文件夹名称来找
	router.GET("/goods/list", func(context *gin.Context) {
		context.HTML(http.StatusOK, "goods/list.html", gin.H{
			"name": "微服务开发",
		})
	})

	router.GET("/users/list", func(context *gin.Context) {
		context.HTML(http.StatusOK, "users/list.html", gin.H{
			"name": "微服务开发",
		})
	})

	router.Run(":8083")
}

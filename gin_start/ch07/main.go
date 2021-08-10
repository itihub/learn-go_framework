package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required,min=3,max=10"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type SignUpForm struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required,min=3"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // 跨字段校验
}

// 去除校验错误ValidationErrors中struct的name
func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

// 初始化校验翻译器
func InitTrans(local string) (err error) {
	// 修改Gin框架中validator引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 注册获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New()              // 中文翻译器
		enT := en.New()              // 英文翻译器
		uni := ut.New(enT, zhT, enT) // 第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		trans, ok = uni.GetTranslator(local)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", local)
		}

		switch local {
		case "en":
			en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			en_translations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

var trans ut.Translator // 全局翻译器

func main() {

	if err := InitTrans("zh"); err != nil {
		fmt.Println("初始化翻译器错误")
		return
	}

	router := gin.Default()

	router.POST("/loginJSON", func(context *gin.Context) {
		// 参数校验绑定
		var loginForm Login
		if err := context.ShouldBind(&loginForm); err != nil {
			errs, ok := err.(validator.ValidationErrors) // 判断是否校验错误
			if !ok {
				context.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})
			}
			context.JSON(http.StatusBadRequest, gin.H{
				"error": removeTopStruct(errs.Translate(trans)), // 翻译错误
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
		})
	})

	router.POST("/signup", func(context *gin.Context) {
		var signUpForm SignUpForm
		if err := context.ShouldBind(&signUpForm); err != nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})

	router.Run(":8083")
}

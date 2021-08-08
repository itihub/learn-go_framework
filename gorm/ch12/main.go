package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// User 拥有并属于多种 language，`user_languages` 是连接表
type User struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"` // 声明多对多关系
}

type Language struct {
	gorm.Model
	Name string
}

func main() {

	// 使用gorm连接到数据库

	// 设置全局的的logger, 作用：执行每个sql语句的时候会打印每一行sql
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢sql阈值
			LogLevel:                  logger.Info, // Log级别
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, // logger全局模式
	})
	if err != nil {
		panic(err)
	}

	//定义一个表结构，将表结构直接生成对应的表 - migrations
	// 迁移 schema
	//_ = db.AutoMigrate(&User{}) // 多对多时AutoMigrate会自动生成中间表以及关联的两个表

	// 插入数据
	//languages := []Language{}
	//languages = append(languages, Language{Name: "go"})
	//languages = append(languages, Language{Name: "java"})
	//user := User{
	//	Languages: languages,
	//}
	//db.Create(&user)

	// 查询 一对多查询
	//var user User
	//db.Preload("Languages").First(&user)
	//for _, languag := range user.Languages {
	//	fmt.Println(languag.Name)
	//}

	//如果我已经取出一个用户来了,但是这个用户我们之前没有使用Preload来加载对应的 Languages
	//不是说用户有Languages我们就一定要取出来
	var user User
	db.First(&user)
	var languages = []Language{}
	db.Model(&user).Association("Languages").Find(&languages)
	for _, language := range languages {
		fmt.Println(language.Name)
	}
}

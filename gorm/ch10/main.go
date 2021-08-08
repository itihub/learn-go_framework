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

// `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	Name      string
	CompanyID int // 数据库中存储的字段company_id
	Company   Company
}

type Company struct {
	ID   int
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
	//_ = db.AutoMigrate(&User{}) // 新建了users表和companys表并设置了外键

	var user User
	db.First(&user)
	fmt.Println(user.Name, user.Company.Name)

	// 使用 Perload、Joins 预加载 belongs to 关联的记录
	db.Preload("Company").First(&user) // Preload 设置 struct 中的外键struct变量名称
	fmt.Println(user.Name, user.Company.Name)

	db.Joins("Company").First(&user)
	fmt.Println(user.Name, user.Company.Name)
}

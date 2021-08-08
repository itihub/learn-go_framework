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

// User 有多张 CreditCard，UserID 是外键
type User struct {
	gorm.Model
	CreditCards []CreditCard `gorm:"foreignKey:UserRefer"` // 定义逻辑外键关系
}

type CreditCard struct {
	gorm.Model
	Number    string
	UserRefer uint
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
	//_ = db.AutoMigrate(&User{})
	//_ = db.AutoMigrate(&CreditCard{})

	/*
		外键约束会让你的数据很完整，即使业务代码考虑不严谨
		在大型系统、高并发系统中不建议使用外键约束，自己在业务层面保证数据一致性
	*/

	// 创建数据
	//user := User{}
	//db.Create(&user)
	//db.Create(&CreditCard{
	//	Number: "12",
	//	UserRefer: user.ID,
	//})
	//db.Create(&CreditCard{
	//	Number: "34",
	//	UserRefer: user.ID,
	//})

	// 通过逻辑外键 进行has many一对多查询
	var queryUser User
	db.Preload("CreditCards").First(&queryUser)
	for _, card := range queryUser.CreditCards {
		fmt.Println(card.Number)
	}
}

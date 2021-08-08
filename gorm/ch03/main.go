package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time     // 解决不能更新零值的问题 指针方式
	MemberNumber sql.NullString // 解决不能更新零值的问题 NullString 方式
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
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
	//_ = db.AutoMigrate(&User{}) // 自动生成表结构

	// Create 新增数据

	// 使用临时对象创建数据 无法获取创建后数据ID
	//db.Create(&User{
	//	Name: "jimmy",
	//})

	// 使用对象创建数据 可以获取创建后数据ID
	user := User{
		Name: "jimmy2",
	}
	fmt.Println(user.ID)
	result := db.Create(&user)
	fmt.Println(user.ID)             //返回插入数据的主键
	fmt.Println(result.Error)        // 返回error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数

	// updates语句不会更新零值，但是update语句可以更零值
	db.Model(&User{ID: 1}).Updates(User{Name: ""})
	db.Model(&User{ID: 1}).Update("Name", "")

	/*
		updates语句 解决更新非零值字段方法
		方法一：设置字段类型为指针
		方法二：使用sql.NullString 类型来解决
	*/
	// 方式一
	empty := ""
	db.Model(&User{ID: 1}).Updates(User{Email: &empty})
	// 方式二
	db.Model(&User{ID: 1}).Updates(User{MemberNumber: sql.NullString{"", true}})

}

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

type NewUser struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time     // 解决不能更新零值的问题 指针方式
	MemberNumber sql.NullString // 解决不能更新零值的问题 NullString 方式
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
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
	_ = db.AutoMigrate(&NewUser{}) // 自动生成表结构

	// 新增数据
	// 批量插入
	var insertUsers = []NewUser{{Name: "jinzhu1", Age: 20}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(&insertUsers)

	for _, user := range insertUsers {
		fmt.Println(user.ID) // 1,2,3
	}

	// 根据主键逻辑删除
	db.Delete(&NewUser{}, 1)

	var users []NewUser
	db.Find(&users)
	for _, user := range users {
		fmt.Println(user.ID)
	}

	// 查找逻辑删除数据
	db.Unscoped().Where("age = 0").Find(&users)

	// 物理删除
	db.Unscoped().Delete(&NewUser{}, 1)

}

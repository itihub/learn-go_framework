package main

import (
	"database/sql"
	"errors"
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

	// 批量插入
	var insertUsers = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(&insertUsers)

	for _, user := range insertUsers {
		fmt.Println(user.ID) // 1,2,3
	}

	// 查询单个数据
	var user User
	// 获取第一条记录（主键升序）
	db.Find(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;
	fmt.Println(user.ID)

	// 获取一条记录，没有指定排序字段
	db.Take(&user)
	// SELECT * FROM users LIMIT 1;

	// 获取最后一条记录（主键降序）
	db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	// 错误处理
	//result := db.First(&user)
	//result.RowsAffected // 返回找到的记录数
	//result.Error        // returns error

	// 检查 是否为 ErrRecordNotFound 错误
	//errors.Is(result.Error, gorm.ErrRecordNotFound)

	// 通过主键查询

	result := db.First(&user, 10)
	// SELECT * FROM users WHERE id = 10;
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("未找到")
	}

	db.First(&user, "10") // 避免sql注入
	// SELECT * FROM users WHERE id = 10;

	var users []User
	db.Find(&users, []int{1, 2, 3})
	// SELECT * FROM users WHERE id IN (1,2,3);

	// 检索全部对象
	result = db.Find(&users)
	fmt.Println("总记录:", result.RowsAffected)
	for _, user := range users {
		fmt.Println(user.ID)
	}

}

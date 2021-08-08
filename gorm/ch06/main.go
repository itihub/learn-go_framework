package main

import (
	"database/sql"
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

	// 通过Where查询

	var user User
	var users []User

	// 获取第一条匹配的记录

	// string 方式
	db.Where("name = ?", "jinzhu").First(&user) //缺点：需要清楚记得数据库字段名称 优点：灵活性高

	/*
		Struct 方式
		通过Struct作为查询条件 无需知道数据库字段名称通过对象属性完成
		优点：不需要了解细节
		缺点；灵活性低且零值不能作为条件
	*/
	db.Where(&User{
		Name: "jinzhu",
	}).Find(&user) //
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

	// 例如 如果您的字段值为 0、''、false 或其他 零值，该字段不会被用于构建查询条件
	db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu";

	/*
		Map 方式
		优点：灵活性高且零值可以被构建查询条件
		缺点：需要知道数据库表字段名称
	*/
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// 获取全部匹配的记录
	db.Where("name <> ?", "jinzhu").Find(&users)
	// SELECT * FROM users WHERE name <> 'jinzhu';

	// IN
	db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

	// LIKE
	db.Where("name LIKE ?", "%jin%").Find(&users)
	// SELECT * FROM users WHERE name LIKE '%jin%';

	// AND
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

	lastWeek := "2000-01-01 00:00:00"
	today := "2000-01-08 00:00:00"

	// Time
	db.Where("updated_at > ?", lastWeek).Find(&users)
	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

	// BETWEEN
	db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

}

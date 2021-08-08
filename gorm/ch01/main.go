package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Product struct {
	gorm.Model // 引用公用字段
	Code       string
	//Code sql.NullString // 解决string默认值“”不更新问题
	Price uint
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
	_ = db.AutoMigrate(&Product{}) // 自动生成表结构

	// Create 新增数据
	db.Create(&Product{Code: "D42", Price: 100})

	// Read 读取数据
	var product Product
	db.First(&product, 1)                 // 根据整形主键查找第一条数据
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的第一条记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段 （也就是go内建数据类型的默认值不进行更新 如 int 默认值 0 string 默认 ""）
	//db.Model(&product).Updates(Product{Price: 200, Code: sql.NullString{"", true}}) // 通过NullString类型解决零值不能被更新问题
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product 并没有执行delete语句，而是逻辑删除
	db.Delete(&product, 1)

}

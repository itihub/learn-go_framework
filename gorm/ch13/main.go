package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

/*
1. 自定义表名
2. 统一给表名添加前缀
*/
type Language struct {
	gorm.Model
	Name    string
	AddTime time.Time // 每个记录创建的时候自动加上当前时间到AddTime中
	//AddTime sql.NullTime // 每个记录创建的时候自动加上当前时间到AddTime中
}

func (l *Language) BeforeCreate(tx *gorm.DB) (err error) {
	l.AddTime = time.Now()
	return
}

//在gorm中可以通过给某一个 struct添加 TableName方法来自定义表名
//func (Language) TableName() string {
//	return "my_language"
//}

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
		// NamingStrategy 与 TableName 不能同时配置，同时配置以 TableName 为准
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "per_", // 统一添加表前缀名
		},
		Logger: newLogger, // logger全局模式
	})
	if err != nil {
		panic(err)
	}

	//定义一个表结构，将表结构直接生成对应的表 - migrations
	// 迁移 schema
	_ = db.AutoMigrate(&Language{})

	db.Create(&Language{Name: "java"})
}

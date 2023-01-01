package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// Product 定义表结构
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:654321@tcp(192.168.137.134:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	// 设置全局的 logger，这个 logger 在我们执行每个 sql 语句的时候会打印每一样 sql
	// sql 才是最重要的，尽量看到每个 api 背后的 sql 语句是什么
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	// 全局模式

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// 定义一个表结构 将表结构直接生成对应的表 - migrations
	// 迁移 schema
	err = db.AutoMigrate(&Product{}) // 此处应该有建表语句
	if err != nil {
		return
	}
}

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

// User 定义表结构
type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
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
	err = db.AutoMigrate(&User{}) // 此处应该有建表语句
	if err != nil {
		panic(err)
	}
	//db.Create(&User{Name: "winnie"})
	//db.Model(&User{ID: 1}).Update("Name", "")
	db.Model(&User{ID: 1}).Updates(User{Name: ""}) //  UPDATE `users` SET `updated_at`='2023-01-02 14:51:55.691' WHERE `id` = 1
	empty := ""
	db.Model(&User{ID: 1}).Updates(User{Email: &empty}) // UPDATE `users` SET `email`='',`updated_at`='2023-01-02 14:51:
	/*
		Updates 语句不会更新零值，Update 语句会更新
		解决仅更新非零值字段的方法有两种
		1.将 string 值设置为 *string
		2.使用 sql 的 NULLxxx 来解决
	*/
}
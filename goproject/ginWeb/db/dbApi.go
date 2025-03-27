package db

import (
	"fmt"
	"time"

	"ginWeb/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dbHandler struct {
	db *gorm.DB
}

type Condition struct {
	Field    string
	Operator string // "=", ">", "<", "LIKE", "IN"...
	Value    interface{}
}

var DbHandler *dbHandler = &dbHandler{}

func InitDB() {
	// MySQL连接示例
	dsn := fmt.Sprintf(
		"%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.Cfg.DB.UserName,
		config.Cfg.DB.Password, config.Cfg.DB.Protocol, config.Cfg.DB.Host, config.Cfg.DB.Port,
		config.Cfg.DB.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database; err: %s", err.Error()))
	}

	// 获取底层的 sql.DB 对象
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("failed to get sql.DB; err: %s", err.Error()))
	}

	// 配置连接池
	sqlDB.SetMaxIdleConns(10)           // 空闲连接池中的最大连接数
	sqlDB.SetMaxOpenConns(100)          // 数据库的最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接可复用的最大时间

	DbHandler.db = db
}

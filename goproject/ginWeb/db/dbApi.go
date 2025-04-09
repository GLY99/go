package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"ginWeb/config"
	"ginWeb/logger"

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

type TransactionFunc func(*gorm.DB) error

var DbHandler *dbHandler = &dbHandler{}

func InitDB() error {
	// MySQL连接示例
	dsn := fmt.Sprintf(
		"%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.Cfg.DB.UserName,
		config.Cfg.DB.Password, config.Cfg.DB.Protocol, config.Cfg.DB.Host, config.Cfg.DB.Port,
		config.Cfg.DB.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database; err: %s", err.Error())
	}

	// 获取底层的 sql.DB 对象
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB; err: %s", err.Error())
	}

	// 配置连接池
	sqlDB.SetMaxIdleConns(10)           // 空闲连接池中的最大连接数
	sqlDB.SetMaxOpenConns(100)          // 数据库的最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接可复用的最大时间

	DbHandler.db = db
	return nil
}

func WithRetry(db *gorm.DB, maxRetries int, delay time.Duration) func(TransactionFunc, ...*sql.TxOptions) error {
	return func(txFunc TransactionFunc, opts ...*sql.TxOptions) error {
		var lastErr error

		for i := 0; i < maxRetries; i++ {
			if i > 0 {
				time.Sleep(delay)
			}

			// 这个logger不是请求上下文中的logger，没有requestid等信息
			logger.Logger.Info(fmt.Sprintf("exec txFunc with %d", i+1))
			lastErr = db.Transaction(txFunc, opts...)
			if lastErr == nil {
				return nil
			}

			if !isRetryableError(lastErr) {
				break
			}
		}
		return fmt.Errorf("transaction failed after %d attempts: %w", maxRetries, lastErr)
	}
}

// isRetryableError 判断错误是否可重试
func isRetryableError(err error) bool {
	if err == nil {
		return false
	}

	// 常见的可重试数据库错误
	retryableErrors := []string{
		"deadlock",
		"timeout",
		"connection",
		"busy",
		"locked",
	}

	errStr := strings.ToLower(err.Error())
	for _, keyword := range retryableErrors {
		if strings.Contains(errStr, keyword) {
			return true
		}
	}

	return false
}

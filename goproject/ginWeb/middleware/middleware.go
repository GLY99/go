package middleware

import (
	"fmt"
	"ginWeb/logger"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// 中间件处理顺序
// router.Use(Middleware1(), Middleware2(), Middleware3())
// 请求进入
//     ↓
// Middleware1 前置代码
//     ↓
// Middleware2 前置代码
//     ↓
// Middleware3 前置代码
//     ↓
// c.Next() 执行后续的中间件的前置代码或者路由函数
//     ↓
// Middleware3 后置代码
//     ↓
// Middleware2 后置代码
//     ↓
// Middleware1 后置代码
//     ↓
// 响应返回

func RegisterMiddlewares(r *gin.Engine) {
	// Gin框架自己的日志，记录请求信息，输出到标准输出
	// r.Use(gin.Logger())
	// gin.Recovery()这个中间件当panic发生时,会回复500错误,但是会将错误输出到标准输出，如果想输出到日志需要自定义RecoveryFunc
	// r.Use(gin.Recovery())

	r.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		stack := string(debug.Stack()) // 获取堆栈
		logger.Logger.Error(fmt.Sprintf("Panic: %v\nStack: %s", err, stack))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": 1,
			"error":  "Internal Server Error",
		})
	}))
	// 生成requestID
	r.Use(RequestIDMiddleware())
	// 注册日志中间件
	r.Use(LoggerMiddleware())
}

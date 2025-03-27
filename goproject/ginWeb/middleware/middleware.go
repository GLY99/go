package middleware

import (
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
// 路由处理函数c.Next()
//     ↑
// Middleware3 后置代码
//     ↑
// Middleware2 后置代码
//     ↑
// Middleware1 后置代码
//     ↑
// 响应返回

func RegisterMiddlewares(r *gin.Engine) {
	// Gin框架自己的日志
	r.Use(gin.Logger())
	// 这个中间件当panic发生时,会回复500错误
	r.Use(gin.Recovery())
	// 生成requestID
	r.Use(RequestIDMiddleware())
	// 注册日志中间件
	r.Use(LoggerMiddleware())
}

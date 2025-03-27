package middleware

import (
	"ginWeb/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := c.MustGet("RequestID").(string)

		// 创建请求专用的Logger
		requestLogger := logger.Logger.With(
			zap.String("requestID", requestID),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("clientIp", c.ClientIP()),
		)
		c.Set("logger", requestLogger)

		// 记录请求开始
		requestLogger.Info("request started")

		c.Next()

		// 记录请求完成
		latency := time.Since(start)
		requestLogger.Info("request completed",
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", latency),
			zap.Int("responseSize", c.Writer.Size()),
		)
	}
}

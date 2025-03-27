package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 尝试从请求头获取X-Request-ID
		requestID := c.GetHeader("X-Request-ID")

		// 若请求头未提供，生成新的UUID
		if requestID == "" {
			newUUID, err := uuid.NewUUID()
			if err != nil {
				// 处理极端情况下的UUID生成失败（如系统熵不足）
				// 此处生成简化备用ID，实际生产环境可优化
				requestID = "req-" + fmt.Sprintf("fallback-%d", time.Now().UnixNano())
			} else {
				requestID = "req-" + newUUID.String()
			}
		}

		// 存储到上下文
		c.Set("RequestID", requestID)
		// 设置响应头
		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Next()
	}
}

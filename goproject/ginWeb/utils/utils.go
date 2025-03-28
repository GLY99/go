package utils

import (
	"context"
	"ginWeb/logger"

	"go.uber.org/zap"
)

// 从context获取logger
func LoggerFromContext(ctx context.Context, loggerKey string) *zap.Logger {
	if logger, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return logger
	}
	return logger.Logger
}

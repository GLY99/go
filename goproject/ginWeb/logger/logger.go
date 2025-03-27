package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func InitLogger() {
	// 配置日志输出位置以及小时轮转
	hourlyRotate := &lumberjack.Logger{
		Filename:   "logs/app.log", // 日志文件路径
		MaxSize:    100,            // 每个日志文件最大尺寸(MB)
		MaxBackups: 168,            // 保留7天的日志(24*7)
		MaxAge:     168,            // 保留7天的日志(小时)
		LocalTime:  true,           // 使用本地时间
		Compress:   true,           // 压缩归档旧日志
	}

	// 编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 可读的时间格式

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel) // 可以动态调整日志级别

	// 创建核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器
		zapcore.AddSync(hourlyRotate),         // 输出到按小时切割的文件
		atomicLevel,                           // 日志级别
	)

	// 创建Logger
	Logger = zap.New(core, zap.AddCaller())
}

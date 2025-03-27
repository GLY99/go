package main

import (
	"context"
	"fmt"
	"ginWeb/config"
	"ginWeb/logger"
	"ginWeb/middleware"
	"ginWeb/routing"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"ginWeb/db"
)

func main() {
	// 加载配置文件
	config.LoadConfig()

	// 初始化日志
	logger.InitLogger()

	// 创建 Gin 引擎实例
	r := gin.New()

	// 注册中间件
	middleware.RegisterMiddlewares(r)

	// 创建db连接
	db.InitDB()

	// 定义路由和处理函数
	routing.InitRouting(r)

	// 启动服务（默认监听 0.0.0.0:8080）
	// r.Run(":8851")
	// 配置 HTTP 服务器
	srv := &http.Server{
		Addr:    ":8851",
		Handler: r,
		// 可选配置
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 10 * time.Second,
		// IdleTimeout:  60 * time.Second,
	}

	// 在协程中启动服务器
	go func() {
		logger.Logger.Info(fmt.Sprintf("服务启动，监听端口 %s", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Fatal(fmt.Sprintf("服务器错误: %v", err))
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	defer close(quit)
	// 监听 SIGINT (Ctrl+C) 和 SIGTERM (kill)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 这里会阻塞直到收到信号
	<-quit
	logger.Logger.Info("正在关闭服务器...")

	// 设置优雅关闭的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 执行优雅关闭
	if err := srv.Shutdown(ctx); err != nil {
		logger.Logger.Fatal(fmt.Sprintf("服务器强制关闭: %v\n", err))
	}

	logger.Logger.Info("服务器已退出")
}

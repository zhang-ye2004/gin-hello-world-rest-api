package main

import (
	"gin-hello-world-rest-api/internal/config"
	"gin-hello-world-rest-api/internal/handler"
	"gin-hello-world-rest-api/internal/logger"
	"gin-hello-world-rest-api/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
	// "honnef.co/go/tools/config"
)

func main() {
	// 1. 加载配置
	cfg := config.Load()

	// 2. 初始化日志
	logger.Init(cfg.LogFile)

	// 3. 设置 Gin 模式
	gin.SetMode(cfg.GinMode)

	// 4. 创建路由
	r := gin.Default()

	// 5. 注册全局中间件（所有路由都会记录日志）
	r.Use(middleware.Logger())

	// 6. 注册路由
	r.GET("/", handler.RootHandler(cfg))
	r.GET("/env", handler.EnvHandler(cfg))
	r.GET("/error", handler.ErrorHandler)

	// 7. 启动服务
	log.Printf("服务器启动在端口:%s", cfg.Port)
	r.Run(":" + cfg.Port)
}

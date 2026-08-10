package handler

import (
	"log"
	"net/http"

	"gin-hello-world-rest-api/internal/config"

	"github.com/gin-gonic/gin"
)

// 根路径处理器
func RootHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Gin!",
			"env":     cfg.GinMode,
		})
	}
}

// 环境变量处理器
func EnvHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"PORT":        cfg.Port,
			"GIN_MODE":    cfg.GinMode,
			"DB_HOST":     cfg.DBHost,
			"DB_PORT":     cfg.DBPort,
			"DB_USER":     cfg.DBUser,
			"DB_PASSWORD": cfg.DBPass,
			"DB_NAME":     cfg.DBName,
		})
	}
}

// 错误测试处理器
func ErrorHandler(c *gin.Context) {
	log.Println("这是一个错误日志测试")
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "something went wrong",
	})
}

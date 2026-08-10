package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// ===== 1. 配置日志输出到文件 =====
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("打开日志文件失败：", err)
	}
	defer logFile.Close()

	// 同时输出到文件和控制台
	multiWriter := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(multiWriter)

	// 设置日志格式（带时间戳和文件名）
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// ===== 2. 读取环境变量 =====
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = "debug"
	}
	gin.SetMode(mode)

	// ===== 3. 启动服务 =====
	r := gin.Default()

	// 中间件：记录每个请求的耗时
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		log.Printf("[%s] &s %s %d %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP,
			c.Writer.Status(),
			latency,
		)
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Heloo from gin!",
			"env":     mode,
		})
	})

	r.GET("/env", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"PORT":     port,
			"GIN_MODE": mode,
		})
	})

	//模拟错误端口（用于测试日志）
	r.GET("/error", func(c *gin.Context) {
		log.Println("这是一个错误日志测试")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
	})

	log.Printf("服务器启动在端口 %s", port)
	r.Run(":" + port)
}

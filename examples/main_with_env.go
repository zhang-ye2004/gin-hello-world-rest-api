package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// 加载 .env 文件（如果存在）
	if err := godotenv.Load(); err != nil {
		log.Println("没有找到 .env 文件，使用系统环境变量")
	}

	// 1. 读取环境变量（带默认值）
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = "debug" //"release"会关闭调试日志，提升性能
	}
	gin.SetMode(mode)

	// 2. 读取数据库连接（示例，这里只是展示）
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// 打印配置（实际部署时，密码不要打印）
	fmt.Println("===== 当前配置 =====")
	fmt.Printf("PORT: %s\n", port)
	fmt.Printf("GIN_MODE: %s\n", mode)
	fmt.Printf("DB_HOST: %s\n", dbHost)
	fmt.Printf("DB_PORT: %s\n", dbPort)
	fmt.Printf("DB_USER: %s\n", dbUser)
	fmt.Printf("DB_NAME: %s\n", dbName)
	fmt.Println("====================")

	// 3. 启动服务
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello from gin!!",
			"env":     mode,
		})
	})

	r.GET("/env", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"PORT":        port,
			"GIN_MODE":    mode,
			"DB_HOST":     dbHost,
			"DB_PORT":     dbPort,
			"DB_USER":     dbUser,
			"DB_PASSWORD": dbPassword,
			"DB_NAME":     dbName,
		})
	})

	log.Println("服务器启动在端口", port)
	r.Run(":" + port)

}

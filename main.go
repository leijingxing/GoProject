package main

import (
	"fmt"
	"go-crud-app/config"
	"go-crud-app/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 加载.env文件
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 初始化数据库连接
	config.InitDB()

	// 创建Gin引擎
	r := gin.Default()

	// 设置路由
	routes.SetupRoutes(r)

	// 获取服务器端口
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	// 启动服务器
	fmt.Printf("Server running on port %s\n", port)
	r.Run(":" + port)
} 
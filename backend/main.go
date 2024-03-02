package main

import (
	"backend/api" // 替換 "your_project_name" 為你的專案名

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化資料庫連接
	api.InitDB()

	// 創建Gin引擎
	router := gin.Default()

	// 定義路由
	router.POST("/api/user/login", api.UserLogin)

	// 啟動服務
	router.Run(":8080")
}

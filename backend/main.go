package main

import (
	"backend/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 设置用户相关的路由
	r.POST("/api/user/register", api.UserRegister)
	r.POST("/api/user/login", api.UserLogin)
	// 启动服务器
	r.Run(":8080")
}

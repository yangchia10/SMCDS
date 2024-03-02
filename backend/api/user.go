package api

import (
    "github.com/gin-gonic/gin"
)

// 用户注册处理函数
func UserRegister(c *gin.Context) {
    // 处理用户注册逻辑...
    c.JSON(200, gin.H{
        "message": "注册成功",
    })
}

// 用户登录处理函数
func UserLogin(c *gin.Context) {
    // 处理用户登录逻辑...
    c.JSON(200, gin.H{
        "message": "登录成功",
    })
}

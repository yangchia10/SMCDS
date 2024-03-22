package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

// UserRegistrationInfo 用於接收註冊信息的結構體
type UserRegistrationInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// UserRegister 處理用戶註冊請求的函數
func UserRegister(c *gin.Context) {
	// 從請求中獲取註冊信息
	var regInfo UserRegistrationInfo
	if err := c.BindJSON(&regInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的請求格式"})
		return
	}

	// 在這裡，你可以添加密碼加密或驗證邏輯

	// 向資料庫中插入新用戶
	_, err := DB.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", regInfo.Username, regInfo.Password, regInfo.Email)
	if err != nil {
		if sqlErr, ok := err.(*mysql.MySQLError); ok && sqlErr.Number == 1062 {
			c.JSON(http.StatusConflict, gin.H{"error": "用戶名已存在"})
		} else {
			log.Printf("資料庫插入錯誤: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "伺服器內部錯誤"})
		}
		return
	}

	// 返回成功響應
	c.JSON(http.StatusCreated, gin.H{"message": "註冊成功"})
}

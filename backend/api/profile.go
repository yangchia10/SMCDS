package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserUpdateInfo 用於接收更新信息的結構體
type UserUpdateInfo struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

// UserUpdate 處理用戶資料更新請求的函數
func UserUpdate(c *gin.Context) {
	// 從URL參數中獲取用戶ID
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的用戶ID"})
		return
	}

	// 從請求中獲取更新信息
	var updateInfo UserUpdateInfo
	if err := c.BindJSON(&updateInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的請求格式"})
		return
	}

	// 建立一個用於存儲 SQL 更新語句和參數的變數
	var updateQuery string
	var queryParams []interface{}

	// // 檢查是否需要更新密碼(密碼加密版)
	// if updateInfo.Password != "" {
	// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateInfo.Password), bcrypt.DefaultCost)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "密碼加密失敗"})
	// 		return
	// 	}
	// 	updateQuery = "UPDATE users SET password = ?"
	// 	queryParams = append(queryParams, hashedPassword)
	// }

	// 檢查是否需要更新密碼
	if updateInfo.Password != "" {
		updateQuery = "UPDATE users SET password = ?"
		queryParams = append(queryParams, updateInfo.Password)
	}

	// 檢查是否需要更新信箱
	if updateInfo.Email != "" {
		if updateQuery != "" {
			updateQuery += ", "
		} else {
			updateQuery = "UPDATE users SET "
		}
		updateQuery += "email = ?"
		queryParams = append(queryParams, updateInfo.Email)
	}

	// 確保有需要更新的字段
	if updateQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "沒有提供更新的資料"})
		return
	}

	// 添加 WHERE 條件到 SQL 語句
	updateQuery += " WHERE id = ?"
	queryParams = append(queryParams, userId)

	// 執行更新操作
	_, err = DB.Exec(updateQuery, queryParams...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新用戶資料失敗"})
		return
	}

	// 返回成功響應
	c.JSON(http.StatusOK, gin.H{"message": "用戶資料更新成功"})
}

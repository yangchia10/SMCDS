package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck 檢查系統健康狀態
func HealthCheck(c *gin.Context) {
	// 嘗試連接數據庫
	if err := DB.Ping(); err != nil {
		// 如果連接失敗，返回錯誤響應
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "unhealthy",
			"message": "無法連接到數據庫",
		})
		return
	}

	// 如果連接成功，返回健康狀態
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

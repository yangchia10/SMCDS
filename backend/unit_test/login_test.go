package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/api" // 替換 "your_project/backend" 為你的專案路徑

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserLogin(t *testing.T) {
	// 初始化 Gin 引擎
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/user/login", api.UserLogin)

	// 創建一個模擬的登錄請求
	loginData := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	jsonData, _ := json.Marshal(loginData)
	req, _ := http.NewRequest("POST", "/api/user/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// 執行測試請求
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 檢查響應狀態碼
	assert.Equal(t, http.StatusOK, w.Code)

	// 檢查響應體
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, loginData["username"], response["username"])
	assert.Equal(t, "user_role", response["role"]) // 替換 "user_role" 為預期的角色
}

package api_test

import (
	"backend/api" // 替換為你的專案的實際路徑
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestUserLogin 測試用戶登錄函數
func TestUserLogin(t *testing.T) {
	// 創建 sqlmock 對象
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("創建 sqlmock 對象時發生錯誤: %s", err)
	}
	defer db.Close()

	// 設置預期的數據庫操作和結果
	rows := sqlmock.NewRows([]string{"id", "username", "role"}).
		AddRow(1, "testuser", "user_role")
	mock.ExpectQuery("^SELECT id, username, role FROM users WHERE username = \\? AND password = \\?$").
		WithArgs("testuser", "testpassword").
		WillReturnRows(rows)

	// 設置 API 的數據庫連接為 mock 的數據庫連接
	api.DB = db

	// 初始化 Gin 並注冊路由
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/user/login", api.UserLogin)

	// 構建請求數據和請求對象
	loginData := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	jsonData, _ := json.Marshal(loginData)
	req, _ := http.NewRequest("POST", "/api/user/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// 執行請求
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 校驗預期的 sqlmock 調用是否全部被滿足
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("存在未滿足的預期: %s", err)
	}

	// 校驗響應狀態碼
	assert.Equal(t, http.StatusOK, w.Code)

	// 打印響應體
	respBody := w.Body.String()
	t.Logf("響應體: %s", respBody)

	// 校驗響應數據
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("解析響應體時發生錯誤: %s", err)
	} else {
		// 打印出 API 響應中的用戶資訊
		t.Logf("用戶 ID: %v", response["id"])
		t.Logf("用戶名: %s", response["username"])
		t.Logf("用戶角色: %s", response["role"])
	}

	// 這裡我們檢查用戶名和角色是否如預期返回
	assert.Equal(t, loginData["username"], response["username"])
	assert.Equal(t, "user_role", response["role"]) // 確保這個值與你的 mock 數據匹配
}

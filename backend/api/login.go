package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
    // 從請求中獲取用戶名和密碼
    var loginInfo struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.BindJSON(&loginInfo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "無效的請求格式"})
        return
    }

    // 在資料庫中查找用戶
    var user struct {
        ID       int    `json:"id"`
        Username string `json:"username"`
        Role     string `json:"role"`
    }
    err := DB.QueryRow("SELECT id, username, role FROM users WHERE username = ? AND password = ?", loginInfo.Username, loginInfo.Password).Scan(&user.ID, &user.Username, &user.Role)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "用戶名或密碼錯誤"})
        } else {
            log.Println("資料庫查詢錯誤:", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "伺服器內部錯誤"})
        }
        return
    }

    // 返回用戶資訊
    c.JSON(http.StatusOK, user)
}

package api_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"backend/api"
)

func TestInitDB(t *testing.T) {
	// 創建 sqlmock 對象
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("創建 sqlmock 對象時發生錯誤: %s", err)
	}
	defer db.Close()

	// 設置預期的數據庫操作和結果
	// mock.ExpectOpen()
	// mock.ExpectPing()

	// 調用 InitDB 函數
	api.DB = db
	api.InitDB()

	// 校驗預期的 sqlmock 調用是否全部被滿足
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("存在未滿足的預期: %s", err)
	}

	// 使用 assert 來驗證 DB 是否成功初始化
	assert.NotNil(t, api.DB, "數據庫連接應該被成功初始化")
}

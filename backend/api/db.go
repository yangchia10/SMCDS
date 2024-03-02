package api

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:1010@tcp(localhost:3306)/smcds?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("無法連接到資料庫: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("無法連接到資料庫: %v", err)
	}

	log.Println("成功連接到資料庫")
}

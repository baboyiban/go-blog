package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	// 연결 문자열 수정
	dsn := "gouser:gopassword@tcp(localhost:3306)/goblog?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

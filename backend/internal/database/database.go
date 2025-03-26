package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	// 환경 변수에서 데이터베이스 연결 정보 가져오기
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// 기본값 설정 (환경 변수가 비어 있을 경우)
	if dbHost == "" {
		dbHost = "mysql:3306" // 기본값은 Docker 네트워크의 서비스 이름
	}

	// DSN 구성
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbName)

	// TLS 설정이 필요한 경우 추가 (옵션)
	if os.Getenv("DB_USE_TLS") == "true" {
		dsn += "&tls=true"
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	// 연결 테스트
	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("데이터베이스 연결 실패: %v", err))
	}

	return db
}

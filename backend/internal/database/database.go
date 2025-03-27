package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	// 환경 변수에서 데이터베이스 연결 정보 가져오기
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbDATABASE := os.Getenv("DB_DATABASE")

	// DSN 구성
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbDATABASE)

	// TLS 설정이 필요한 경우 추가 (옵션)
	if os.Getenv("DB_USE_TLS") == "true" {
		dsn += "&tls=true"
	}

	var err error

	retryCount := 10              // 재시도 횟수 증가
	retryDelay := 5 * time.Second // 딜레이 증가

	for i := range retryCount {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("시도 %d: 연결 실패 - %v", i+1, err)
			time.Sleep(retryDelay)
			continue
		}

		// 연결 풀 설정 추가
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
		db.SetConnMaxLifetime(5 * time.Minute)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := db.PingContext(ctx); err != nil { // Context 사용
			log.Printf("시도 %d: 핑 실패 - %v", i+1, err)
			db.Close()
			time.Sleep(retryDelay)
			continue
		}

		log.Println("✅ 데이터베이스 연결 성공")
		return db
	}

	panic(fmt.Sprintf("연결 실패(%d회 재시도): %v", retryCount, err))
}

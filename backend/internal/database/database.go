package database

import (
	"database/sql"
	"fmt"
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

	var db *sql.DB
	var err error
	retryCount := 5               // 재시도 횟수 설정
	retryDelay := 3 * time.Second // 재시도 간 딜레이 설정

	for i := range retryCount {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			fmt.Printf("데이터베이스 연결 시도 %d 실패 (sql.Open 에러): %v\n", i+1, err)
			time.Sleep(retryDelay)
			continue
		}

		err = db.Ping()
		if err != nil {
			fmt.Printf("데이터베이스 연결 시도 %d 실패 (db.Ping 에러): %v\n", i+1, err)
			if db != nil {
				db.Close() // Ping 실패 시 연결 닫기 (자원 누수 방지)
			}
			time.Sleep(retryDelay)
			continue
		}

		fmt.Println("데이터베이스 연결 성공!")
		return db // 연결 성공 시 DB 객체 반환
	}

	panic(fmt.Sprintf("데이터베이스 연결 실패 (최대 재시도 횟수 초과): %v", err)) // 최대 재시도 횟수 초과 시 패닉
}

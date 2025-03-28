package database

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName: os.Getenv("DB_NAME"),
		Params: map[string]string{
			"tls":     os.Getenv("DB_SSL_MODE"),
			"timeout": "30s", // 연결 타임아웃
		},
		AllowNativePasswords: true,
		Collation:            "utf8mb4_unicode_ci",
	}

	// SSL 인증서 로드
	if os.Getenv("DB_SSL_MODE") != "disabled" {
		rootCertPool := x509.NewCertPool()
		pem, _ := os.ReadFile(os.Getenv("DB_CA_CERT"))
		if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
			return nil, errors.New("DB SSL 인증서 불러오기 실패")
		}
		mysql.RegisterTLSConfig("custom", &tls.Config{
			RootCAs:    rootCertPool,
			MinVersion: tls.VersionTLS12,
		})
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("DB 연결 초기화 실패: %v", err)
	}

	// 연결 풀 설정
	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetMaxOpenConns(50) // 외부 DB는 풀 크기 증가
	db.SetMaxIdleConns(25)

	return db, nil
}

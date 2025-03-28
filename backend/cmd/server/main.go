package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/baboyiban/go-blog/internal/database"
	"github.com/baboyiban/go-blog/internal/routes"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// 데이터베이스 연결
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf(`
		치명적 오류: 데이터베이스 연결 실패
		호스트: %s
		에러: %v
		`, os.Getenv("MYSQL_HOST"), err)
	}
	defer db.Close()

	// 라우터 설정
	router := routes.SetupRouter(db)

	// 서버 설정
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("서버 시작 - 포트 %s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("서버 종료: %v", err)
	}
}

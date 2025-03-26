package main

import (
	"log"
	"net/http"
	"os"

	"github.com/baboyiban/go-blog/internal/database"
	"github.com/baboyiban/go-blog/internal/routes"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
    db := database.ConnectDB()
    defer db.Close()

    router := routes.SetupRouter(db)

    // 환경변수에서 포트 가져오기
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // 기본값
    }

    log.Printf("Starting server on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}

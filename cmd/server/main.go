package main

import (
	"go-blog/internal/database"
	"go-blog/internal/handlers"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	router := handlers.NewRouter(db)
	log.Fatal(http.ListenAndServe(":8080", router))
}

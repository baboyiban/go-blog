package main

import (
	"log"
	"net/http"

	"github.com/baboyiban/go-blog/internal/database"
	"github.com/baboyiban/go-blog/internal/handlers"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	router := handlers.NewRouter(db)
	log.Fatal(http.ListenAndServe(":8080", router))
}

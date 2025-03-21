package main

import (
	"log"
	"net/http"

	"github.com/baboyiban/go-blog/internal/database"
	"github.com/baboyiban/go-blog/internal/routes"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	router := routes.SetupRouter(db) // Now matches the signature
	log.Fatal(http.ListenAndServe(":8080", router))
}

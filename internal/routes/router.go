package routes

import (
	"database/sql"

	"github.com/baboyiban/go-blog/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine { // Add db parameter
	r := gin.Default()

	userHandler := handlers.NewUserHandler(db)

	// Correct route definition for Gin
	r.GET("/users/:id", userHandler.GetUser)

	return r
}

package routes

import (
	"github.com/baboyiban/go-blog/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/users", handlers.GetUsers)
    return r
}

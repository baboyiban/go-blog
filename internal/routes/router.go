package routes

import (
    "github.com/gin-gonic/gin"
    "go-blog/internal/handlers"  // 경로 수정
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/users", handlers.GetUsers)
    return r
}

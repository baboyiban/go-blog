package routes

import (
	"database/sql"

	"github.com/baboyiban/go-blog/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	// 1. Svelte 빌드 결과물 서빙
	router.Static("/static", "./web/static")

	// 2. SPA 라우팅 처리 (모든 미등록 경로는 index.html로)
	router.NoRoute(func(c *gin.Context) {
		c.File("./web/static/index.html")
	})

	// 3. API 라우트 설정
	router.GET("/api/posts", handlers.GetPosts)

	return router
}

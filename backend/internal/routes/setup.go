package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	// 1. API 라우트 설정
	api := router.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
		// 다른 API 엔드포인트들...
	}

	// 2. SvelteKit 애플리케이션을 위한 프록시 설정
	// 개발 환경
	if gin.Mode() != gin.ReleaseMode {
		// Vite dev 서버로 프록시
		router.Any("/*path", func(c *gin.Context) {
			// Vite dev 서버로 프록시
			// TODO: 프록시 구현
		})
	} else {
		// 프로덕션 환경
		// SvelteKit의 빌드된 SSR 핸들러로 요청 전달
		router.Static("/static", "./web/static")
		router.NoRoute(func(c *gin.Context) {
			// TODO: SvelteKit SSR 핸들러 연동
		})
	}

	return router
}

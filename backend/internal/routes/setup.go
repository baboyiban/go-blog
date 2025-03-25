package routes

import (
	"database/sql"
	"net/http/httputil"
	"net/url"
	"strings"

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

	// 2. 개발/프로덕션 모드 분기
	if gin.Mode() != gin.ReleaseMode {
		setupDevProxy(router) // 개발 모드 프록시 설정
	} else {
		setupProductionRoutes(router) // 프로덕션 라우트 설정
	}

	return router
}

func setupDevProxy(router *gin.Engine) {
	target, _ := url.Parse("http://frontend:3000") // 프론트엔드 포트 일치 확인
	proxy := httputil.NewSingleHostReverseProxy(target)

	router.NoRoute(func(c *gin.Context) {
		// 헤더 설정 추가
		c.Request.Header.Set("X-Forwarded-Host", c.Request.Host)
		proxy.ServeHTTP(c.Writer, c.Request)
	})
}

func setupProductionRoutes(router *gin.Engine) {
	// 정적 파일 서빙 설정
	router.Static("/static", "./web/static")

	// SPA Fallback 처리
	router.NoRoute(func(c *gin.Context) {
		c.File("./web/index.html")
	})
}

func isAPIRequest(path string) bool {
	return strings.HasPrefix(path, "/api")
}

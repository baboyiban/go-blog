package routes

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	// HTML 템플릿 자동 리로드 설정
	router.HTMLRender = &render.HTMLDebug{
		Glob: "web/templates/*.html",
	}

	// 정적 파일 서빙 설정 추가
	router.Static("/static", "web/static")

	// 기존 라우트 설정...
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":   "Go Blog",
			"Message": "Welcome to Go Blog!",
		})
	})

	return router
}

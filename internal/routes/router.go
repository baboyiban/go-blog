package routes

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// 템플릿 로드 설정 추가
	r.LoadHTMLGlob("web/templates/*")

	// 정적 파일 서빙 설정 추가
	r.Static("/static", "./web/static")

	// 기존 라우트 설정...
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":   "Go Blog",
			"Message": "Welcome to Go Blog!",
		})
	})

	return r
}

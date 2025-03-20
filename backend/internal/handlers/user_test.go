package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/baboyiban/go-blog/internal/handlers"
	"github.com/gin-gonic/gin"
)

func TestGetUser(t *testing.T) {
	// Create Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	// Create handler and test
	handler := handlers.NewUserHandler(nil) // Pass nil or mock DB
	handler.GetUser(c)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

package handlers_test  // 테스트 전용 패키지

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "go-blog/internal/handlers"
)

func TestGetUser(t *testing.T) {
	req := httptest.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()

	handler := UserHandler{}
	handler.GetUser(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

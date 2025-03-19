package handlers

import (
	"database/sql" // 추가
	"encoding/json"
	"net/http"

	"github.com/baboyiban/go-blog/internal/models"
)

type UserHandler struct {
	DB *sql.DB
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// DB 조회 로직
	user := models.User{ID: 1, Name: "John"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

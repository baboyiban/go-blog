package handlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{db: db}
}

// Gin-compatible handler signature
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	// Your implementation using gin.Context
	c.JSON(200, gin.H{"id": id})
}

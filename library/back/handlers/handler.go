package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	db_pool *pgxpool.Pool
}

func New_handler(db_pool *pgxpool.Pool) *Handler {
	return &Handler{db_pool: db_pool}
}

func (h *Handler) Status_handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "online",
		"message": "Conexión a DB exitosa y servidor corriendo",
	})
}

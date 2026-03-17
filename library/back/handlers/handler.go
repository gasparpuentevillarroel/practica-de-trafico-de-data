package handlers

import (
	"net/http"
	"strconv"

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

func (h *Handler) Get_book_handler(c *gin.Context) {
	book_id := c.Param("id")

	if _, err := strconv.Atoi(book_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	// Placeholder mientras se implementa consulta real a DB.
	c.JSON(http.StatusOK, gin.H{
		"id":      book_id,
		"message": "book encontrado",
	})
}

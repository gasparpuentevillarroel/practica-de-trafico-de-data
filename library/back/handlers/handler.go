package handlers

import (
	"library/back/db"
	"library/back/models"
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

func (h *Handler) Add_book(c *gin.Context) {
	type book_input struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		Autor_name string `json:"autor_name"`
		Autor_id   int64  `json:"autor_id"`
		Year       int64  `json:"year"`
	}

	var req book_input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
		return
	}

	// Valida el libro
	newBook, err := models.New_book(req.Id, req.Name, req.Autor_name, req.Autor_id, req.Year)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Una sola consulta INSERT (con constraint UNIQUE en la BD)
	if err := db.Insert_book(c.Request.Context(), h.db_pool, newBook); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         newBook.Id_value(),
		"name":       newBook.Name_value(),
		"autor_name": newBook.Autor_value(),
		"autor_id":   newBook.Autor_id_value(),
		"year":       newBook.Year_value(),
		"created_at": newBook.Created_at_value(),
		"updated_at": newBook.Updated_at_value(),
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

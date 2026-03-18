package handlers

import (
	"errors"
	"library/back/db"
	"library/back/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type book_input struct {
	Id               string `json:"id"`
	Title            string `json:"title"`
	Author_name      string `json:"author_name"`
	Author_id        int64  `json:"author_id"`
	Year_publication int64  `json:"year_publication"`
}

func (h *Handler) Add_book(c *gin.Context) {
	//convierto datos de entrada en tipo book_input
	var req book_input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
		return
	}

	//lo valido con el constructor
	newBook, err := models.New_book(req.Id, req.Title, req.Author_name, req.Author_id, req.Year_publication)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Insert_book(c.Request.Context(), h.db_pool, newBook); err != nil {
		if errors.Is(err, db.Err_book_already_exists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               newBook.Id_value(),
		"title":            newBook.Title_value(),
		"author_name":      newBook.Author_name_value(),
		"author_id":        newBook.Author_id_value(),
		"year_publication": newBook.Year_publication_value(),
		"created_at":       newBook.Created_at_value(),
		"updated_at":       newBook.Updated_at_value(),
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

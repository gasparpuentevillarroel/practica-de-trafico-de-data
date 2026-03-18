package handlers

import (
	"errors"
	"library/back/db"
	"library/back/models"
	"library/back/security"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user_input struct {
	id       string `json:user_id`
	name     string `json: user_name`
	password string `json:user_password`
}

func (h *Handler) Add_user(c *gin.Context) {

	var req user_input
	//convierto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
		return
	}
	//hasheo
	if req.password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
	}
	secure_password, err := security.HashPassword(req.password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	//valido
	new_user, err := models.New_user(req.id, req.name, secure_password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := db.Insert_user(c.Request.Context(), h.db_pool, *new_user); err != nil {
		if errors.Is(err, db.Err_user_already_exists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user_id":   new_user.Id(),
		"user_name": new_user.Name(),
		"pasword":   req.password,
	})

}

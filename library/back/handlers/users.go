package handlers

import (
	"errors"
	"library/back/db"
	"library/back/models"
	"library/back/security"
	"net/http"

	"github.com/gin-gonic/gin"
)

type signup_input struct {
	Id       string `json:"user_id"`
	Name     string `json:"user_name"`
	Password string `json:"user_password"`
}

type login_input struct {
	Id       string `json:"user_id"`
	Password string `json:"user_password"`
}

func (h *Handler) Add_user(c *gin.Context) {

	var req signup_input
	//convierto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
		return
	}
	//hasheo
	if req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}
	secure_password, err := security.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//valido
	new_user, err := models.New_user(req.Id, req.Name, secure_password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
		"pasword":   req.Password,
	})

}

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var req login_input

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
		return
	}

	if req.Id == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id and user_password are required"})
		return
	}
	user, err := db.WhereUser(ctx, h.db_pool, req.Id)
	if err != nil {
		if errors.Is(err, db.Err_user_not_found) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
	}

	is_valide := security.CheckPaswordHash(req.Password, user.Password())
	if !is_valide {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "contraseña incorrecta"})
		return
	} else {
		token, err := security.GenerateJWT(user.Id())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo generar el acceso"})
			return
		}

		// 200 OK: Le enviamos el token a React
		c.JSON(http.StatusOK, gin.H{
			"message": "login exitoso",
			"token":   token,
			"user_id": user.Id(),
		})
	}
}

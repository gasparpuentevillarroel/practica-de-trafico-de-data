package main

import (
	"library/back/handlers"
	"library/back/security"

	"github.com/gin-gonic/gin"
)

func register_routes(router *gin.Engine, handler_instance *handlers.Handler) {
	// --- RUTAS PÚBLICAS (Sin protección) ---
	router.GET("/status", handler_instance.Status_handler)

	router.POST("/register", handler_instance.Add_user)
	router.POST("/login", handler_instance.Login)

	// --- RUTAS PROTEGIDAS (Requieren Token RSA) ---
	protected := router.Group("/api")
	protected.Use(security.AuthMiddleware())
	{
		protected.GET("/books/:id", handler_instance.Get_book_handler)
		protected.POST("/books", handler_instance.Add_book)
	}
}

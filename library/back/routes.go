package main

import (
	"practicar/back/handlers"

	"github.com/gin-gonic/gin"
)

func register_routes(router *gin.Engine, handler_instance *handlers.Handler) {
	router.GET("/status", handler_instance.Status_handler)
	router.GET("/books/:id", handler_instance.Get_book_handler)
}

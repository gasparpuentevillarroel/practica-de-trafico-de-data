package main

import (
	"library/back/handlers"

	"github.com/gin-gonic/gin"
)

func register_routes(router *gin.Engine, handler_instance *handlers.Handler) {
	router.GET("/status", handler_instance.Status_handler)
	router.GET("/books/:id", handler_instance.Get_book_handler)
	router.POST("/books", handler_instance.Add_book)
}

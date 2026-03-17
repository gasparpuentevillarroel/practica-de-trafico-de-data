package main

import (
	"context"
	"fmt"
	"log"
	"os" // Para leer el puerto del .env
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"library/back/db"
	"library/back/handlers"
)

func main() {
	godotenv.Load()

	conn_str := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db_pool, err := db.Connect_db(ctx, conn_str)
	if err != nil {
		log.Fatal(err)
	}
	defer db_pool.Close()

	// 3. Configurar Gin
	http_router := gin.Default()
	handler_instance := handlers.New_handler(db_pool)
	register_routes(http_router, handler_instance)

	// 5. Levantar el servidor
	app_port := os.Getenv("PORT")
	if app_port == "" {
		app_port = "8080" // Puerto por defecto si no está en el .env
	}

	log.Printf("Servidor arrancando en el puerto %s...", app_port)
	if err := http_router.Run(":" + app_port); err != nil {
		log.Fatalf("No se pudo arrancar el servidor: %v", err)
	}
}

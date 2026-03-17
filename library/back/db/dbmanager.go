package db

import (
	"context"
	"errors"
	"fmt"
	"practicar/back/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type db_manager struct {
}

func Connect_db(conn_str string) (*pgxpool.Pool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db_pool, err := pgxpool.New(ctx, conn_str)
	if err != nil {
		return nil, fmt.Errorf("no se pudo crear el pool: %w", err)
	}

	if err := db_pool.Ping(ctx); err != nil {
		db_pool.Close()
		return nil, fmt.Errorf("error al verificar conexión: %w", err)
	}

	fmt.Println("¡Conexión exitosa a PostgreSQL!")
	return db_pool, nil
}

func get_book_by_id(in_id string) (models.Book, error) {
	_ = in_id
	return models.Book{}, errors.New("get_book_by_id pendiente de implementación")
}

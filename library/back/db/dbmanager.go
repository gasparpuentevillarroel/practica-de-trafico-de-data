package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type db_manager struct {
}

var Err_book_already_exists = errors.New("el libro ya existe (violación de constraint único)")
var Err_user_already_exists = errors.New("el usuario ya existe (violación de constraint único)")

func Connect_db(ctx context.Context, conn_str string) (*pgxpool.Pool, error) {
	db_pool, err := pgxpool.New(ctx, conn_str)
	if err != nil {
		return nil, fmt.Errorf("no se pudo crear el pool: %w", err)
	}

	if err := db_pool.Ping(ctx); err != nil {
		db_pool.Close()
		return nil, fmt.Errorf("error al verificar conexión: %w", err)
	}

	fmt.Println("Conexión exitosa a PostgreSQL")
	return db_pool, nil
}

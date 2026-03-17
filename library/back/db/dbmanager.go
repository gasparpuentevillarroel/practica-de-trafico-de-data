package db

import (
	"context"
	"errors"
	"fmt"
	"library/back/models"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type db_manager struct {
}

func Connect_db(ctx context.Context, conn_str string) (*pgxpool.Pool, error) {
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

func Insert_book(ctx context.Context, db_pool *pgxpool.Pool, book *models.Book) error {
	query := `
		INSERT INTO books (id, name, autor_name, autor_id, year, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := db_pool.Exec(ctx, query,
		book.Id_value(),
		book.Name_value(),
		book.Autor_value(),
		book.Autor_id_value(),
		book.Year_value(),
		book.Created_at_value(),
		book.Updated_at_value(),
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return errors.New("el libro ya existe (violación de constraint único)")
		}
		return fmt.Errorf("error al insertar libro: %w", err)
	}

	return nil
}

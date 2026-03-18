package db

import (
	"context"
	"errors"
	"fmt"
	"library/back/models"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Insert_book(ctx context.Context, db_pool *pgxpool.Pool, book *models.Book) error {
	query := `
		INSERT INTO books (id, title, author_name, author_id, year_publication, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := db_pool.Exec(ctx, query,
		book.Id_value(),
		book.Title_value(),
		book.Author_name_value(),
		book.Author_id_value(),
		book.Year_publication_value(),
		book.Created_at_value(),
		book.Updated_at_value(),
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return Err_book_already_exists
		}
		return fmt.Errorf("error al insertar libro: %w", err)
	}

	return nil
}

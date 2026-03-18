package db

import (
	"context"
	"errors"
	"fmt"
	"library/back/models"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Insert_user(ctx context.Context, db_pool *pgxpool.Pool, in_user models.User) error {
	query := `INSERT INTO users  (id,name,password,created_at)
			VALUES($1,$2,$3,$4)`

	_, err := db_pool.Exec(ctx, query,
		in_user.Id(),
		in_user.Name(),
		in_user.Password(),
		time.Now().UTC(),
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return Err_user_already_exists
		}
		return fmt.Errorf("error al insertar usuario: %w", err)
	}

	return nil
}

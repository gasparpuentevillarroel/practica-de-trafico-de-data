package models

import (
	"fmt"
	"time"
)

type Book struct {
	id         string    `json:"id"`
	name       string    `json:"name"`
	autor_name string    `json:"autor_name"`
	autor_id   int64     `json:"autor_id"`
	year       int64     `json:"year"`
	created_at time.Time `json:"created_at"`
	updated_at time.Time `json:"updated_at"`
}

func New_book(in_id string, in_name string, in_autor string, in_autor_id int64, in_year int64) (*Book, error) {
	if err := validate_string(in_id); err != nil {
		return nil, err
	}
	if err := validate_string(in_name); err != nil {
		return nil, err
	}
	if err := validate_string(in_autor); err != nil {
		return nil, err
	}
	if err := validate_autor_id(in_autor_id); err != nil {
		return nil, err
	}
	if err := validate_year(in_year); err != nil {
		return nil, err
	}

	return &Book{
		id:         in_id,
		name:       in_name,
		autor_name: in_autor,
		autor_id:   in_autor_id,
		year:       in_year,
		created_at: time.Now(),
		updated_at: time.Now(),
	}, nil
}

func (b Book) String() string {
	return fmt.Sprintf("Book<id=%s name=%s autor=%s autor_id=%d>", b.id, b.name, b.autor_name, b.autor_id)
}

// getters
func (bk Book) Id_value() string {
	return bk.id
}

func (bk Book) Name_value() string {
	return bk.name
}

func (bk Book) Autor_value() string {
	return bk.autor_name
}

func (bk Book) Autor_id_value() int64 {
	return bk.autor_id
}

func (bk Book) Year_value() int64 {
	return bk.year
}

func (bk Book) Created_at_value() time.Time {
	return bk.created_at
}

func (bk Book) Updated_at_value() time.Time {
	return bk.updated_at
}

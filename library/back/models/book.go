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

func new_book(in_id string, in_name string, in_autor string, in_year int64) (*Book, error) {
	if err := validate_string(in_id); err != nil {
		return nil, err
	}
	if err := validate_string(in_name); err != nil {
		return nil, err
	}
	if err := validate_string(in_autor); err != nil {
		return nil, err
	}
	if err := validate_year(in_year); err != nil {
		return nil, err
	}

	return &Book{
		id:         in_id,
		name:       in_name,
		autor_name: in_autor,
		year:       in_year,
		created_at: time.Now(),
		updated_at: time.Now(),
	}, nil
}

func (b Book) String() string {
	return fmt.Sprintf("Book<id=%s name=%s autor=%s>", b.id, b.name, b.autor_name)
}

//getters

func (bk Book) id_value() string {
	return bk.id
}

func (bk Book) name_value() string {
	return bk.name
}

func (bk Book) autor_value() string {
	return bk.autor_name
}

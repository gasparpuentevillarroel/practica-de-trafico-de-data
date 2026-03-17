package models

import (
	"fmt"
	"time"
)

type Book struct {
	id               string    `json:"id"`
	title            string    `json:"title"`
	author_name      string    `json:"author_name"`
	author_id        int64     `json:"author_id"`
	year_publication int64     `json:"year_publication"`
	created_at       time.Time `json:"created_at"`
	updated_at       time.Time `json:"updated_at"`
}

func New_book(in_id string, in_title string, in_author string, in_author_id int64, in_year_publication int64) (*Book, error) {
	if err := validate_string(in_id); err != nil {
		return nil, err
	}
	if err := validate_string(in_title); err != nil {
		return nil, err
	}
	if err := validate_string(in_author); err != nil {
		return nil, err
	}
	if err := validate_autor_id(in_author_id); err != nil {
		return nil, err
	}
	if err := validate_year(in_year_publication); err != nil {
		return nil, err
	}

	return &Book{
		id:               in_id,
		title:            in_title,
		author_name:      in_author,
		author_id:        in_author_id,
		year_publication: in_year_publication,
		created_at:       time.Now(),
		updated_at:       time.Now(),
	}, nil
}

func (b Book) String() string {
	return fmt.Sprintf("Book<id=%s title=%s author=%s author_id=%d>", b.id, b.title, b.author_name, b.author_id)
}

// getters
func (bk Book) Id_value() string {
	return bk.id
}

func (bk Book) Title_value() string {
	return bk.title
}

func (bk Book) Author_name_value() string {
	return bk.author_name
}

func (bk Book) Author_id_value() int64 {
	return bk.author_id
}

func (bk Book) Year_publication_value() int64 {
	return bk.year_publication
}

func (bk Book) Created_at_value() time.Time {
	return bk.created_at
}

func (bk Book) Updated_at_value() time.Time {
	return bk.updated_at
}

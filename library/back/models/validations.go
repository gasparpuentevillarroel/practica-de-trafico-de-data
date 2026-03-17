package models

import (
	"errors"
)

func validate_string(in_string string) error {
	if in_string == "" {
		return errors.New("el string debe estar completo")
	}
	return nil
}

func validate_autor_id(in_autor_id int64) error {
	if in_autor_id <= 0 {
		return errors.New("autor_id debe ser mayor a 0")
	}
	return nil
}

func validate_year(in_year int64) error {
	if in_year < 1000 || in_year > 9999 {
		return errors.New("el año debe tener formato YYYY")
	}
	return nil
}

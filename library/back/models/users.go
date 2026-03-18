package models

import "time"

type User struct {
	id         string
	name       string
	password   string
	created_at time.Time
}

func New_user(in_id string, in_name string, in_password string) (*User, error) {

	if err := validate_string(in_id); err != nil {
		return nil, err
	}
	if err := validate_string(in_name); err != nil {
		return nil, err
	}
	if err := validate_string(in_password); err != nil {
		return nil, err
	}

	nw_user := User{
		id:         in_id,
		name:       in_name,
		password:   in_password,
		created_at: time.Now().UTC(),
	}

	return &nw_user, nil
}

//getters

func (u User) Id() string {
	return u.id
}
func (u User) Name() string {
	return u.name
}
func (u User) Password() string {
	return u.password
}

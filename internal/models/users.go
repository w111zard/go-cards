package models

import (
	"database/sql"
)

type User struct {
	ID           string
	Email        string
	PasswordHash string
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Create(email string, password string) error {
	return nil
}

func (m *UserModel) GetAll() ([]*User, error) {
	return nil, nil
}

func (m *UserModel) GetByID(id string) (*User, error) {
	return nil, nil
}

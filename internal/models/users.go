package models

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Create(name string, email string, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	query := "INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3)"
	_, err = m.DB.Exec(query, name, email, hash)

	return err
}

func (m *UserModel) GetAll() ([]*User, error) {
	query := "SELECT id, name, email, password_hash FROM users"

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (m *UserModel) GetByID(id string) (*User, error) {
	query := `SELECT id, name, email, password_hash from "users" where id = $1`

	var u User

	err := m.DB.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &u, nil
}

func (m *UserModel) GetByEmail(email string) (*User, error) {
	query := `SELECT id, name, email, password_hash from "users" where email = $1`

	var u User

	err := m.DB.QueryRow(query, email).Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &u, nil
}

package models

import (
	"database/sql"
	"errors"
)

type Deck struct {
	ID     string
	Name   string
	UserID string
}

type DeckModel struct {
	DB *sql.DB
}

func (m *DeckModel) Create(name string, userID string) error {
	query := `INSERT INTO decks (name, user_id)
	VALUES ($1, $2)`

	if _, err := m.DB.Exec(query, name, userID); err != nil {
		return err
	}

	return nil
}

func (m *DeckModel) GetByID(id string) (*Deck, error) {
	query := `SELECT id, name, user_id from "decks" where id = $1`

	var d Deck

	err := m.DB.QueryRow(query, id).Scan(&d.ID, &d.Name, &d.UserID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &d, nil
}

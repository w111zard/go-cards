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

func (m *DeckModel) GetAllByUserID(userID string) ([]*Deck, error) {
	query := `SELECT id, name, user_id FROM "decks" WHERE user_id = $1`

	rows, err := m.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var decks []*Deck
	for rows.Next() {
		var d Deck
		if err := rows.Scan(&d.ID, &d.Name, &d.UserID); err != nil {
			return nil, err
		}
		decks = append(decks, &d)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return decks, nil
}

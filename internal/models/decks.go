package models

import (
	"database/sql"
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
	return nil
}

func (m *DeckModel) GetByID(id string) (*DeckModel, error) {
	return nil, nil
}

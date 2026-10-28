package main

import (
	"github.com/w111zard/go-cards/internal/models"
)

type templateData struct {
	User  *models.User
	Decks []*models.Deck
}

package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/w111zard/go-cards/internal/models"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
	decks  *models.DeckModel
	users  *models.UserModel
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", "postgres://wizard:1234@localhost:5432/go_cards?sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := connectDB()
	if err != nil {
		logger.Error("db err", err)
	}
	defer db.Close()
	logger.Info("successfully connected to db")

	app := &application{
		logger: logger,
		decks:  &models.DeckModel{db},
		users:  &models.UserModel{db},
	}

	server := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	logger.Info("server is running", "address", *addr)
	err = server.ListenAndServe()
	logger.Error("server error", err)
	os.Exit(1)
}

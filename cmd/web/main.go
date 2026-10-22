package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	server := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	logger.Info("server is running", "address", *addr)
	err := server.ListenAndServe()
	logger.Error("server error", err)
	os.Exit(1)
}

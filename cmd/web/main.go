package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	logger.Info("server is running", "port", 3000)
	err := http.ListenAndServe(":3000", mux)
	logger.Error("server error", err)
	os.Exit(1)
}

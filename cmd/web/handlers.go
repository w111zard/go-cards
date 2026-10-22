package main

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("go-cards home page"))
}

package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) findUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	user, err := app.users.GetByID(id)
	if err != nil {
		app.serverError(w, err)
		return
	}
	if user == nil {
		app.clientError(w, http.StatusNotFound)
		return
	}

	js, err := json.Marshal(user)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

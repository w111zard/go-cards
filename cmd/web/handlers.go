package main

import (
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	userID := "5af92237-164a-4bd5-93f2-db7ab8eb7c51"

	user, err := app.users.GetByID(userID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	decks, err := app.decks.GetAllByUserID(userID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{
		User:  user,
		Decks: decks,
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}

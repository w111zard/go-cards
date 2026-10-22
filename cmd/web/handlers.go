package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

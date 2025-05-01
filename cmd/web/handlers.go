package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
)

type application struct {
	logger *slog.Logger
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string) {
	// Slice containing paths to HTMLs. Base template must be *first*
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		fmt.Sprintf("./ui/html/pages/%s", t),
	}

	// Unpack strings and process them individually
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Write content of the "base" template as response body
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.tmpl.html")
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "about.tmpl.html")
}

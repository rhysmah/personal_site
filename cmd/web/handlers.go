package main

import (
	"log"
	"net/http"
	"text/template"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/home.html",
		"./ui/html/partials/nav.html",
	}

	// Attempt to read template file into template set (ts)
	ts, err := template.ParseFiles(files...)
	if err != nil {

		// Log the error, then send error for display to the user
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return // return so no subsequent code is run
	}

	// Write template content as response body
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func resumeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/pdf")
	http.ServeFile(w, r, "./ui/static/Rhys_Mahannah_Resume_Aug_2024.pdf")
}

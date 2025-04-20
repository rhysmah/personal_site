package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Slice containing paths to HTMLs
	// Base template must be *first*
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	// Read the files and store templates into template set
	// '...' is used for variadic arguments -- it unpacks a slice of strings
	// and passes each string in the slice as an argument to be processed.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return // return so no subsequent code is executed
	}

	// Write content of the "base" template as response body
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

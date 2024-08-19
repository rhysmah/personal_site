package main

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hi, my name is Rhys"))
}

func resumeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/pdf")
	http.ServeFile(w, r, "./ui/static/Rhys_Mahannah_Resume_Aug_2024.pdf")
}

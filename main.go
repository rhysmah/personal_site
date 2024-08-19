package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, my name is Rhys Mahannah."))
}

func resumeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Rhys_Mahannah_Resume_Aug_2024.pdf")
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", homeHandler)
	mux.HandleFunc("GET /resume", resumeHandler)

	log.Println("starting on server :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

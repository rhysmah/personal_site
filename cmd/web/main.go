package main

import (
	"log"
	"net/http"
)

// Home handler writes a byte slice containing text
// This text is the response BODY
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /about", about)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

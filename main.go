package main

import (
	"log"
	"net/http"
)

// Home handler writes a byte slice containing text
// This text is the response BODY
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test text"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("starting server on :4000")

	// This starts a new web server. It takes two parameters:
	// - TCP network address to listen on
	// - mux, where the requests received by the server will be passed
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

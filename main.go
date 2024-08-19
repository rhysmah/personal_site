package main

import (
	"log"
	"net/http"
)

// A home handler function.
// Writes a byte slice containing a string to the response body.
// Write() uses a byte slice because different types of data may be written
// all data is composed fundamentally of bytes, so using a byte slice gives
// us the flexibility to write different data types.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, my name is Rhys Mahannah."))
}

func resumeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Rhys_Mahannah_Resume_Aug_2024.pdf")
}

func main() {

	// Initialize a new server mux; we can then register
	// a new home function that will be called whenever
	// we navigate to the "/" address; this is a catch-all,
	// meaning that, right now, all requests will be handled
	// by this handler.
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/resume", resumeHandler)

	log.Println("starting on server :4000")

	// Starts a new web server. This takes two arguments:
	// the address where requests are received and served;
	// and the servemux, which handles the requests.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, my name is Rhys Mahannah"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("starting on server :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

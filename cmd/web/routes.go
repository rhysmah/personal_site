package main

import "net/http"

// routes returns a multiplexer
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /about", app.about)

	return mux
}

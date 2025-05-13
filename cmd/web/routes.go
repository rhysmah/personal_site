package main

import "net/http"

// routes returns a multiplexer
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /about", app.about)

	// middleware
	handler := app.secureHeaders(mux)
	handler = app.logRequest(handler)
	handler = app.recoverPanic(handler)

	return handler
}

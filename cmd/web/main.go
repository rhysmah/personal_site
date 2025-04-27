package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/rhysmah/personal_site/internal/log"
)

func main() {

	// CLI-based configuration
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Instantiate custom logger
	logger := log.Default()

	// Instantiate custom application type; inject the custom logger.
	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /about", app.about)

	// .Any() prevents !BADKEY errors
	app.logger.Info("Starting server", slog.Any("addr", *addr))

	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		app.logger.Error("Server failed", slog.String("error", err.Error()), slog.String("addr", *addr))
		os.Exit(1)
	}
}

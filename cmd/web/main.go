package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {

	// CLI-based configuration
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Structured Logger
	// Setting minimum level to LevelDebug, which are the least severe
	// and are, by default, silently discarded from logs
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true, // includes filename and log number
	}))

	// Instantiate the struct that'll use the logger
	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /about", app.about)

	// .Any() prevents !BADKEY errors
	logger.Info("Starting server", slog.Any("addr", *addr))

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error()) // Human-readable message
	os.Exit(1)                // No equivalent to log.Fatal(), which calls this automatically
}

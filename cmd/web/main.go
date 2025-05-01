package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	// CLI-based configuration
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Instantiate handler options and logger for injection
	handlerOpts := slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &handlerOpts))

	app := &application{
		logger: logger,
	}

	// .Any() prevents !BADKEY errors
	app.logger.Info("Starting server", slog.Any("addr", *addr))

	err := http.ListenAndServe(*addr, app.routes())
	if err != nil {
		app.logger.Error("Server failed", slog.String("error", err.Error()), slog.String("addr", *addr))
		os.Exit(1)
	}
}

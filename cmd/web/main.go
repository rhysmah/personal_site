package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

type SiteData struct {
	HomeHeader string `json:"homeHeader"`
	HomeBody   string `json:"homeBody"`
}

type application struct {
	logger    *slog.Logger
	templates map[string]*template.Template
	siteData  SiteData
}

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
		logger:    logger,
		templates: make(map[string]*template.Template),
	}

	// Load data from JSON file
	data, err := os.ReadFile("./data/data.json")
	if err != nil {
		app.logger.Error("Failed to read JSON data file", "error", err)
		os.Exit(1)
	}

	if err := json.Unmarshal(data, &app.siteData); err != nil {
		app.logger.Error("Failed to parse JSON data file", "error", err)
		os.Exit(1)
	}

	// .Any() prevents !BADKEY errors
	app.logger.Info("Starting server", slog.Any("addr", *addr))

	err = http.ListenAndServe(*addr, app.routes())
	if err != nil {
		app.logger.Error("Server failed", slog.String("error", err.Error()), slog.String("addr", *addr))
		os.Exit(1)
	}
}

package log

import (
	"log/slog"
	"os"
)

// Logger wraps slog.Loggers to create an application-specific custom logger
// This uses *slog.Logger as an embedded type, an "is-a" relationship, meaning
// the customer Logger type has automatic access to all the slog methods (versus
// the named type, a "has-a" relationship, which requires manual forwarding).
type Logger struct {
	*slog.Logger
}

func New(level slog.Level, addSource bool) *Logger {
	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: addSource,
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	return &Logger{slog.New(handler)}
}

func Default() *Logger {
	return New(slog.LevelDebug, true)
}

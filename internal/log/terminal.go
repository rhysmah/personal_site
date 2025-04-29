package log

import (
	"context"
	"io"
	"log/slog"
	"sync"
)

type TerminalHandler struct {
	// Where logs are written
	out io.Writer

	// Minimum log level to process
	minLevel slog.Level

	//Default attributes to be included with every log message.
	attrs slog.Attr

	//
	mu sync.Mutex
}

func NewTerminalHandler(w io.Writer, opts *slog.HandlerOptions) *TerminalHandler {
	// If no options, use defaults
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}

	return &TerminalHandler{
		out:      w,
		minLevel: opts.Level.Level(),
	}
}

// Enabled reports whether the handler handles records at the given level.
func (h *TerminalHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.minLevel
}

// Handle processes a log record and formats it for terminal output
func (h *TerminalHandler) Handle(_ context.Context, r slog.Record) error {

}

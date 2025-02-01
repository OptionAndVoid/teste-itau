package logging

import (
	"io"
	"log/slog"
)

func SetDefaultJSONLogger(w io.Writer, opts *slog.HandlerOptions) {
	jsonHandler := slog.NewJSONHandler(w, opts)
	logger := slog.New(jsonHandler)

	slog.SetDefault(logger)
}

package main

import (
	"log/slog"
	"os"
)

// arquive for examples
func main() {
	// h := slog.NewTextHandler(os.Stderr, nil)
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	log := slog.New(h).With("app", "exp")

	log.Debug("Debug message")
	log.Info("Info message", slog.Group("request_info", "request_id", 123))
	log.Warn("warn message")
	log.Error("error message", "request_id", 123)
}

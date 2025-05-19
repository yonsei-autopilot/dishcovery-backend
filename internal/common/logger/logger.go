package logger

import (
	"log/slog"
	"os"
)

func InitializeLogger() {
	handlerOpts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stderr, handlerOpts))
	slog.SetDefault(logger)
}

package util

import (
	"log"
	"log/slog"
	"os"
	"strings"
)

func SetupLogger() {
	logFormat := strings.ToLower(os.Getenv("LOG_FORMAT"))

	var logLevel slog.Level
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		logLevel = slog.LevelDebug
	case "info", "":
		logLevel = slog.LevelInfo
	case "warning", "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		log.Fatalf("Unknown log level: %s", os.Getenv("LOG_LEVEL"))
	}

	var logger *slog.Logger
	opts := slog.HandlerOptions{Level: logLevel}
	if logFormat == "json" {
		logger = slog.New(slog.NewJSONHandler(os.Stderr, &opts))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stderr, &opts))
	}

	slog.SetDefault(logger)
}

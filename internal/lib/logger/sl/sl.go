package sl

import (
	envConst "KidneySmartBackend/internal/env"
	"log"
	"os"

	"golang.org/x/exp/slog"
)

func SetupLogger(env string, logLevel string) *slog.Logger {
	var logger *slog.Logger
	level := parseLogLevel(logLevel)

	switch env {
	case envConst.Local:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	case envConst.Dev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	case envConst.Prod:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	default:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		logger.Warn("Unknown environment: %s, using default logger settings", env)
	}

	return logger
}
func Err(err error) slog.Attr {

	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func parseLogLevel(level string) slog.Level {
	if !isValidLogLevel(level) {
		log.Fatalf("Invalid logging level: %s", level)
	}

	switch level {
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	case "debug":
		return slog.LevelDebug
	default:
		return slog.LevelInfo
	}
}
func isValidLogLevel(level string) bool {
	switch level {
	case "debug", "info", "warn", "error":
		return true
	default:
		return false
	}
}

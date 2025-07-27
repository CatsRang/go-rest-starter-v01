package util

import (
	"log/slog"
	"os"
	"strings"
	"sync"
)

var (
	logger     *slog.Logger
	loggerOnce sync.Once
)

func InitLogger(level string) {
	loggerOnce.Do(func() {
		logLevel := parseLogLevel(level)
		
		opts := &slog.HandlerOptions{
			Level: logLevel,
		}
		
		handler := slog.NewJSONHandler(os.Stdout, opts)
		logger = slog.New(handler)
		slog.SetDefault(logger)
	})
}

func parseLogLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func GetLogger() *slog.Logger {
	if logger == nil {
		InitLogger("info")
	}
	return logger
}
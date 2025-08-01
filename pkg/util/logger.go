package util

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	logger     *slog.Logger
	loggerOnce sync.Once
)

func parseLogLevel(levelStr string) slog.Level {
	switch strings.ToLower(levelStr) {
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

// InitLogger initializes and sets the global logger with the given log level string
func InitLogger(levelStr string) {
	loggerOnce.Do(func() {
		level := parseLogLevel(levelStr)

		opts := &slog.HandlerOptions{
			Level:     level,
			AddSource: true,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.SourceKey {
					source := a.Value.Any().(*slog.Source)
					return slog.Group("source",
						slog.String("file", filepath.Base(source.File)),
						slog.Int("line", source.Line),
					)
				}
				return a
			},
		}

		handler := slog.NewJSONHandler(os.Stdout, opts)
		logger = slog.New(handler)
		slog.SetDefault(logger)
	})
}

// GetLogger returns the initialized logger instance
func GetLogger() *slog.Logger {
	if logger == nil {
		InitLogger("info")
	}
	return logger
}

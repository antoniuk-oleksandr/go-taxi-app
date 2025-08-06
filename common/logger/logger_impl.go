package logger

import (
	"log/slog"
	"os"
)

type logger struct{
	slogLogger *slog.Logger
}

func NewLogger() Logger {
	return &logger{
		slogLogger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}

func (logger *logger) Error(value ...any) {
	panic("unimplemented")
}

func (logger *logger) Info(msg string, args ...any) {
	logger.slogLogger.Info(msg, args)
}

package logger

type Logger interface {
	Info(msg string, args ...any)
	Error(value ...any)
}

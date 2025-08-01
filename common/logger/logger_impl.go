package logger

import "log"

type logger struct{}

func (l *logger) Log(value any) {
	log.Println(value)
}

func NewLogger() Logger {
	return &logger{}
}

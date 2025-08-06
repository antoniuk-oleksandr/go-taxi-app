package main

import "github.com/antoniuk-oleksandr/go-taxi-app/common/logger"

func main() {
	logger := logger.NewLogger()
	logger.Info("Hello, world!")
}

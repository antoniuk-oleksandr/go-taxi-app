package main

import "common/logger"

func main() {
	logger := logger.NewLogger()
	logger.Log("test")
}

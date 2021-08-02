package logging

import (
	"testing"
)

func TestConsoleLogger(t *testing.T) {
	logger := ConsoleLogger{}
	logger.NewConsoleLogger("logging:TestConsoleLogger")
	t.Run("Logs an info message", func(t *testing.T) {
		logger.Info("This is an info message")
	})

	t.Run("Logs an error message", func(t *testing.T) {
		logger.Error("This is an error message")
	})

	t.Run("Logs a debug message", func(t *testing.T) {
		logger.Debug("This is a debug message")
	})

	t.Run("Logs an warning message", func(t *testing.T) {
		logger.Warn("This is a warning message")
	})
}

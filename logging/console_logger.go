package logging

import (
	"log"
	"os"

	"github.com/gokhantamkoc/growth-framework/utility/console"
)

type ConsoleLogger struct {
	LoggingConfig Config
	Logger        *log.Logger
}

func (dl *ConsoleLogger) InitializeLogger() {
	dl.Logger = log.Default()
}

func (dl *ConsoleLogger) Warn(message string) {
	message = console.ColorYellow + "WARN: " + message + console.ColorReset
	dl.Logger.SetOutput(os.Stdout)
	dl.Logger.Println(message)
}

func (dl *ConsoleLogger) Error(message string) {
	message = console.ColorRed + "ERROR: " + message + console.ColorReset
	dl.Logger.SetOutput(os.Stderr)
	dl.Logger.Println(message)
}

func (dl *ConsoleLogger) Info(message string) {
	message = console.ColorBlue + "INFO: " + message + console.ColorReset
	dl.Logger.SetOutput(os.Stdout)
	dl.Logger.Println(message)
}

func (dl *ConsoleLogger) Debug(message string) {
	message = console.ColorGreen + "DEBUG: " + message + console.ColorReset
	dl.Logger.SetOutput(os.Stdout)
	dl.Logger.Println(message)
}

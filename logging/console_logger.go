package logging

import (
	"fmt"
	"log"
	"os"

	"github.com/gokhantamkoc/growth-framework/utility/console"
)

type ConsoleLogger struct {
	PackageFunctionName string
	LoggingConfig       Config
	Logger              *log.Logger
}

func (dl *ConsoleLogger) NewConsoleLogger(pkgFuncName string) {
	dl.Logger = log.Default()
	dl.PackageFunctionName = pkgFuncName
}

func (dl *ConsoleLogger) Warn(message string) {
	message = fmt.Sprintf(
		"%s %s %s %s %s",
		console.ColorYellow,
		dl.PackageFunctionName,
		WarningHead,
		message,
		console.ColorReset,
		)
	dl.Logger.SetOutput(os.Stdout)
	dl.Logger.Println(message)
}

func (dl *ConsoleLogger) Error(message string) {
	message = fmt.Sprintf(
		"%s %s %s %s %s",
		console.ColorRed,
		dl.PackageFunctionName,
		ErrorHead,
		message,
		console.ColorReset,
	)
	dl.Logger.SetOutput(os.Stderr)
	dl.Logger.Println(message)
}

func (dl *ConsoleLogger) Info(message string) {
	message = fmt.Sprintf(
		"%s %s %s %s %s",
		console.ColorBlue,
		dl.PackageFunctionName,
		InfoHead,
		message,
		console.ColorReset,
	)
	dl.Logger.SetOutput(os.Stdout)
	dl.Logger.Println(message)
}

func (dl *ConsoleLogger) Debug(message string) {
	message = fmt.Sprintf(
		"%s %s %s %s %s",
		console.ColorGreen,
		dl.PackageFunctionName,
		DebugHead,
		message,
		console.ColorReset,
	)
	dl.Logger.SetOutput(os.Stdout)
	dl.Logger.Println(message)
}

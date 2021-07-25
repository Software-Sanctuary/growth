package logging

type Config struct {
	LogLevel string
}

type FileLoggingConfig struct {
	Config
	FilePath string
}

type Type int

const (
	FileLogging Type = iota
	ConsoleLogging
)

package logging

type Logger interface {
	Warn(message string)
	Error(message string)
	Info(message string)
	Debug(message string)
}

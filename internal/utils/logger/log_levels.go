package logger

type LogLevel string

const (
	LogLevelInfo  LogLevel = "info"
	LogLevelError LogLevel = "error"
	LogLevelFatal LogLevel = "fatal"
	LogLevelWarn  LogLevel = "warn"
)

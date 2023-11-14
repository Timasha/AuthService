package logger

import "io"

type Logger interface {
	Log(logMsg LogMsg) error
	LogTrace(traceLogMsg TraceLogMsg) error
	AddSideLogger(writer io.Writer) error
}

package logger

import "time"

type LogMsg struct {
	Time     time.Time
	LogLevel LogLevel
	Msg      string
}

package logdrivers

import (
	"auth/internal/utils/logger"
	"io"

	"github.com/rs/zerolog"
)

type ZerologDriver struct {
	loggers []zerolog.Logger
}

func NewZerologDriver(writers []io.Writer) (z *ZerologDriver) {
	z = &ZerologDriver{}
	for i := 0; i < len(writers); i++ {
		z.loggers = append(z.loggers, zerolog.New(writers[i]))
	}
	return
}

func (z *ZerologDriver) Log(logMsg logger.LogMsg) error {
	var zerologLoglevel zerolog.Level
	switch logMsg.LogLevel {
	case logger.LogLevelError:
		{
			zerologLoglevel = zerolog.ErrorLevel
		}
	case logger.LogLevelInfo:
		{
			zerologLoglevel = zerolog.InfoLevel
		}
	case logger.LogLevelWarn:
		{
			zerologLoglevel = zerolog.WarnLevel
		}
	case logger.LogLevelFatal:
		{
			zerologLoglevel = zerolog.FatalLevel
		}
	}
	for i := 0; i < len(z.loggers); i++ {

		z.loggers[i].WithLevel(zerologLoglevel).Msg(logMsg.Msg)
	}
	return nil
}
func (z *ZerologDriver) LogTrace(traceLogMsg logger.TraceLogMsg) error {
	//TODO: make log trace
	return nil
}
func (z *ZerologDriver) AddSideLogger(writer io.Writer) error {
	z.loggers = append(z.loggers, zerolog.New(writer))
	return nil
}

package pkg

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

type Logger struct {
	file *os.File
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) SetLog(logDirname string, logFilename string) {
	err := os.MkdirAll(logDirname, os.ModePerm)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(logDirname+"/"+logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	l.file = file

	multi := zerolog.MultiLevelWriter(os.Stdout, l.file)
	logger := zerolog.New(multi).With().Timestamp().Logger()

	mode := os.Getenv("GO_ENV")

	if mode == "DEV" || mode == "DEBUG" || mode == "TEST" {
		logger = logger.Level(zerolog.DebugLevel)
	} else {
		logger = logger.Level(zerolog.InfoLevel)
	}

	log.Logger = logger

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

}

func (l *Logger) Close() error {
	return l.file.Close()
}

func (l *Logger) ErrorLog(err error, message string) {
	log.Error().Stack().Err(err).Msg(message)
}

func (l *Logger) PanicLog(err error, message string) {
	log.Panic().Stack().Err(err).Msg(message)
}

func (l *Logger) WarnLog(message string) {
	log.Warn().Stack().Msg(message)
}

func (l *Logger) InfoLog(message string) {
	log.Info().Stack().Msg(message)
}

func (l *Logger) DebugLog(message string) {
	log.Debug().Msg(message)
}

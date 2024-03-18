package common

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var logFile *os.File

func SetLog(logDirname string, logFilename string) {
	err := os.MkdirAll(logDirname, os.ModePerm)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(logDirname+"/"+logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	logFile = file

	multi := zerolog.MultiLevelWriter(os.Stdout, logFile)
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

func LogFileCloser() error {
	return logFile.Close()
}
func ErrorLog(err error, message string) {
	log.Error().Stack().Err(err).Msg(message)
}

func PanicLog(err error, message string) {
	log.Panic().Stack().Err(err).Msg(message)
}

func WarnLog(message string) {
	log.Warn().Stack().Msg(message)
}

func InfoLog(message string) {
	log.Info().Stack().Msg(message)
}

func DebugLog(message string) {
	log.Debug().Msg(message)
}

package log

import (
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Stamp})
}

func Info(msg string) {
	log.Info().Caller(1).Msg(msg)
}

func Infof(format string, v ...interface{}) {
	log.Info().Caller(1).Msgf(format, v...)
}

func Debug(msg string) {
	log.Debug().Caller(1).Msg(msg)
}

func Debugf(format string, v ...interface{}) {
	log.Debug().Caller(1).Msgf(format, v...)
}

func Error(err error) {
	if _, ok := err.(stackTracer); ok {
		log.Error().Caller(1).Msgf("%+v", err)
	} else {
		log.Error().Caller(1).Msg(err.Error())
	}
}

func Fatal(err error) {
	if _, ok := err.(stackTracer); ok {
		log.Fatal().Caller(1).Msgf("%+v", err)
	} else {
		log.Fatal().Caller(1).Msg(err.Error())
	}
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

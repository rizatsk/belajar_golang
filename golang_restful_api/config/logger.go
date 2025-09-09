package config

import (
	"context"
	"golang_restful_api/model/logger"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// inisialisasi base logger global
var baseLogger zerolog.Logger

func LoggerInit() {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	baseLogger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	log.Logger = baseLogger
}

func GetBaseLogger() zerolog.Logger {
	return baseLogger
}

func loggingWithContext(context context.Context, level zerolog.Level, payload logger.Logger) {
	newZeroLog := zerolog.Ctx(context)
	if newZeroLog == nil {
		newZeroLog = &baseLogger
	}

	event := newZeroLog.WithLevel(level).
		Str("module", os.Getenv("SERVICE_NAME")).
		Interface("data", payload.Data).
		Interface("err", payload.Err)

	event.Msg(payload.LogMessage)
}

func logging(level zerolog.Level, payload logger.Logger) {
	event := baseLogger.WithLevel(level).
		Str("module", os.Getenv("SERVICE_NAME")).
		Interface("data", payload.Data).
		Interface("err", payload.Err)

	event.Msg(payload.LogMessage)
}

func LoggerInfo(message string) {
	var dataLog logger.Logger
	dataLog.LogMessage = message
	logging(zerolog.InfoLevel, dataLog)
}

func LoggerInfoWithContext(context context.Context, message string) {
	var dataLog logger.Logger
	dataLog.LogMessage = message
	loggingWithContext(context, zerolog.InfoLevel, dataLog)
}

func LoggerErrorWithContext(context context.Context, error logger.LoggerError) {
	var dataLog logger.Logger
	dataLog.LogMessage = error.Message
	dataLog.Err = error.Error
	loggingWithContext(context, zerolog.ErrorLevel, dataLog)
}

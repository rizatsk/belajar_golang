package helper

import (
	"context"
	"golang_restful_api/config"
	"golang_restful_api/model/logger"
)

func PanicIfError(err error, ctxs ...context.Context) {
	if err != nil {
		if len(ctxs) > 0 && ctxs[0] != nil {
			config.LoggerErrorWithContext(ctxs[0], logger.LoggerError{
				Message: "Reponse error",
				Error:   err,
			})
		}

		panic(err)
	}
}

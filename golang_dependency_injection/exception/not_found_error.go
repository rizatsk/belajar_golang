package exception

import (
	"context"
	"golang_restful_api/helper"
	"golang_restful_api/model/logger"
)

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string, ctxs ...context.Context) NotFoundError {
	if len(ctxs) > 0 && ctxs[0] != nil {
		helper.LoggerErrorWithContext(ctxs[0], logger.LoggerError{
			Message: "Reponse error",
			Error:   error,
		})
	}

	return NotFoundError{Error: error}
}

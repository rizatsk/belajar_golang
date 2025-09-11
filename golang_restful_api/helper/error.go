package helper

import (
	"context"
	"golang_restful_api/config"
	"golang_restful_api/model/logger"
)

type PanicErrorParam struct {
	Err     error
	Message string
	Ctx     context.Context
}

func PanicIfError(param PanicErrorParam) {
	if param.Err != nil {
		if param.Ctx != nil {
			if param.Message == "" {
				param.Message = "Response error"
			}

			config.LoggerErrorWithContext(param.Ctx, logger.LoggerError{
				Message: param.Message,
				Error:   param.Err,
			})
		}

		panic(param.Err)
	}
}

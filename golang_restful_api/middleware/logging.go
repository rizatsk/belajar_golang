package middleware

import (
	"golang_restful_api/config"
	"golang_restful_api/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ctxKey string

const (
	RequestIDKey ctxKey = "request_id"
	PathKey      ctxKey = "path"
)

func LoggingMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
		// generate request_id
		reqID := helper.GenerateUuidV4()

		baseLogger := config.GetBaseLogger()
		reqLog := baseLogger.With().
			Str("request_id", reqID).
			Str("method", request.Method).
			Str("path", request.URL.Path).
			Logger()

		// simpan logger ke ctx
		ctx := reqLog.WithContext(request.Context())

		config.LoggerInfoWithContext(ctx, "Request API")
		next(writer, request.WithContext(ctx), param)
	}
}

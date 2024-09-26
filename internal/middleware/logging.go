package middleware

import (
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func LoggingMiddleware(cfg *config.ServiceConfig, logger logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		// intercepting response status
		wrapped := &responseWriter{ResponseWriter: response, status: http.StatusOK}

		defer func() {
			if err := recover(); err != nil {
				wrapped.WriteHeader(http.StatusInternalServerError)
				logger.Error(
					"error occurred while executing handler",
					slog.Any("err", err),
					slog.String("trace", string(debug.Stack())),
				)
			}
		}()

		start := time.Now()
		next.ServeHTTP(wrapped, request)
		duration := time.Since(start)

		logger.Info(
			"request completed",
			slog.Int("status", wrapped.status),
			slog.String("method", request.Method),
			slog.String("path", request.URL.EscapedPath()),
			slog.Int64("duration_ms", duration.Milliseconds()),
		)
	})
}

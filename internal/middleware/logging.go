package middleware

import (
	"context"
	"net/http"
	"runtime/debug"
	"time"

	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
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

		requestID := uuid.New()
		ctx := context.WithValue(request.Context(), utils.RequestIDKey{}, requestID)

		defer func() {
			if err := recover(); err != nil {
				wrapped.WriteHeader(http.StatusInternalServerError)
				logger.Error(
					"error occurred while executing handler",
					requestID,
					zap.Any("err", err),
					zap.String("trace", string(debug.Stack())),
				)
			}
		}()

		start := time.Now()
		next.ServeHTTP(wrapped, request.WithContext(ctx))
		duration := time.Since(start)

		logger.Info(
			"request completed",
			requestID,
			zap.Int("status", wrapped.status),
			zap.String("method", request.Method),
			zap.String("path", request.URL.EscapedPath()),
			zap.Int64("duration_ms", duration.Milliseconds()),
		)
	})
}

package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/metrics"
)

func MetricsMiddleware(metrics *metrics.Metrics, next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		start := time.Now()

		rec := &statusRecorder{ResponseWriter: response, statusCode: http.StatusOK}

		metrics.ActiveConnections.Inc()
		defer metrics.ActiveConnections.Dec()

		next.ServeHTTP(rec, request)

		duration := time.Since(start).Seconds()
		method := request.Method
		url := request.URL.Path
		status := strconv.Itoa(rec.statusCode)

		metrics.RequestCounter.WithLabelValues(method, url, status).Inc()
		metrics.RequestDuration.WithLabelValues(method, url).Observe(duration)

		responseSize := float64(rec.size)
		metrics.ResponseSizeHistogram.WithLabelValues(method, url).Observe(responseSize)

		if rec.statusCode >= 400 {
			metrics.ErrorCounter.WithLabelValues(method, url, status).Inc()
		}
	})
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
	size       int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

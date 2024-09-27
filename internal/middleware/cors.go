package middleware

import (
	"net/http"
	"os"
)

const (
	allowMethods     = "POST, GET, OPTIONS, PUT, DELETE"
	allowHeaders     = "Content-Type"
	allowCredentials = "true"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowOrigin := os.Getenv("CORS_ORIGIN")

		w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", allowMethods)
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)
		w.Header().Set("Access-Control-Allow-Credentials", allowCredentials)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

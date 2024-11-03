package middleware

import (
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
)

func CORSMiddleware(cfg *config.CORSConfig, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", cfg.AllowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", cfg.AllowMethods)
		w.Header().Set("Access-Control-Allow-Headers", cfg.AllowHeaders)
		w.Header().Set("Access-Control-Expose-Headers", cfg.ExposeHeaders)
		w.Header().Set("Access-Control-Allow-Credentials", fmt.Sprintf("%t", cfg.AllowCredentials))

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

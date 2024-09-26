package middleware

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/jwt"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

func AuthMiddleware(cfg *config.Config, logger logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie(cfg.Auth.Jwt.Cookie.Name)
		if err != nil {
			logger.Warnf("no JWT token found in cookies: %v", err)
			utils.JSONError(response, http.StatusUnauthorized, "unauthorized")
			return
		}

		err = jwt.Verify(cfg, cookie.Value)
		if err != nil {
			logger.Warnf("invalid jwt token: %v", err)
			utils.JSONError(response, http.StatusUnauthorized, "unauthorized")
			return
		}

		next.ServeHTTP(response, request)
	})
}

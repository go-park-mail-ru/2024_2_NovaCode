package middleware

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

func AdminMiddleware(cfg *config.AuthConfig, logger logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie(cfg.Jwt.Cookie.Name)
		if err != nil {
			logger.Warnf("jwt token not found in cookies: %v", err)
			utils.JSONError(response, http.StatusUnauthorized, "unauthorized")
			return
		}

		claims, err := utils.VerifyJWT(&cfg.Jwt, cookie.Value)
		if err != nil {
			logger.Warnf("invalid jwt token: %v", err)
			utils.JSONError(response, http.StatusUnauthorized, "unauthorized")
			return
		}

		role, ok := claims["role"].(string)
		if !ok || role != "admin" {
			logger.Warnf("insufficient permissions: user role is %v", role)
			utils.JSONError(response, http.StatusForbidden, "forbidden: admin access required")
			return
		}

		next.ServeHTTP(response, request)
	})
}

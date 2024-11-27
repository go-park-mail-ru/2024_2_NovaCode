package middleware

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/google/uuid"
)

func AuthMiddleware(cfg *config.AuthConfig, logger logger.Logger, next http.Handler) http.Handler {
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

		userIDStr, ok := claims["user_id"].(string)
		if !ok {
			logger.Warnf("user_id claim not found in token")
			utils.JSONError(response, http.StatusUnauthorized, "unauthorized")
			return
		}

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			logger.Warnf("invalid user ID format in token: %v", err)
			utils.JSONError(response, http.StatusUnauthorized, "unauthorized")
			return
		}

		ctx := context.WithValue(request.Context(), utils.UserIDKey{}, userID)
		next.ServeHTTP(response, request.WithContext(ctx))
	})
}

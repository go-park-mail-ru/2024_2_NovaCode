package middleware

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func AuthMiddleware(cfg *config.AuthConfig, logger logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie(cfg.Jwt.Cookie.Name)
		if err != nil {
			logger.Warnf("jwt token not found in cookies: %v", err)
			utils.JSONError(response, http.StatusUnauthorized, "unauthorized")
			return
		}

		userID, err := utils.VerifyJWT(&cfg.Jwt, cookie.Value)
		if err != nil {
			logger.Warnf("invalid jwt token: %v", err)
			utils.JSONError(response, http.StatusUnauthorized, "unauthorized")
			return
		}

		requestUserIDStr, ok := mux.Vars(request)["user_id"]
		if ok {
			requestUserID, err := uuid.Parse(requestUserIDStr)
			if err != nil {
				logger.Warnf("invalid requested user ID format: %v", err)
				utils.JSONError(response, http.StatusBadRequest, "invalid user id format")
				return
			}

			if userID != requestUserID {
				logger.Warnf("requested user id doesn't match with actual: %v", err)
				utils.JSONError(response, http.StatusForbidden, "forbidden")
				return
			}
		}

		ctx := context.WithValue(request.Context(), utils.UserIDKey{}, userID)
		next.ServeHTTP(response, request.WithContext(ctx))
	})
}

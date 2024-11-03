package middleware

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/csrf"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/google/uuid"
)

func CSRFMiddleware(cfg *config.CSRFConfig, logger logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		token := request.Header.Get(cfg.HeaderName)
		if token == "" {
			logger.Warn("csrf token not provided")
			utils.JSONError(response, http.StatusForbidden, "forbidden")
			return
		}

		userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
		if !ok {
			logger.Errorf("user id not found in context")
			utils.JSONError(response, http.StatusForbidden, "user is not authorized")
			return
		}

		if !csrf.Validate(token, userID.String(), cfg.Salt) {
			logger.Errorf("user id not found in context")
			utils.JSONError(response, http.StatusForbidden, "invalid csrf token")
			return
		}

		next.ServeHTTP(response, request)
	})
}

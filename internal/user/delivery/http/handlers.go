package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/user"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

type userHandlers struct {
	cfg     *config.AuthConfig
	usecase user.Usecase
	logger  logger.Logger
}

func NewUserHandlers(cfg *config.AuthConfig, usecase user.Usecase, logger logger.Logger) user.Handlers {
	return &userHandlers{cfg, usecase, logger}
}

// Health godoc
// @Summary Health check
// @Description Returns "OK" if the service is running
// @Success 200 {string} utils.MessageResponse "OK"
// @Router /api/v1/health [get]
func (handlers *userHandlers) Health(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	if err := utils.WriteResponse(response, http.StatusOK, "OK"); err != nil {
		handlers.logger.Errorf("error encoding health response: %v", err)
		utils.JSONError(response, http.StatusInternalServerError, "failed to get health status")
		return
	}
}

// Register godoc
// @Summary Register new user
// @Description Register new user with username, email and password
// @Accept json
// @Produce json
// @Param user body models.User true "User registration details"
// @Success 200 {object} dto.UserTokenDTO "User registration successful"
// @Failure 400 {object} utils.ErrorResponse "Invalid request body or missing fields"
// @Failure 500 {object} utils.ErrorResponse "Failed to return token"
// @Router /api/v1/auth/register [post]
func (handlers *userHandlers) Register(response http.ResponseWriter, request *http.Request) {
	var user models.User

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		utils.JSONError(response, http.StatusBadRequest, "invalid request body")
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		utils.JSONError(response, http.StatusBadRequest, "username, email and password are required")
		return
	}

	if user.Role == "" {
		user.Role = "regular"
	}

	userTokenDTO, err := handlers.usecase.Register(request.Context(), &user)
	if err != nil {
		handlers.logger.Errorf("failed to register user: %v", err)
		utils.JSONError(response, http.StatusBadRequest, "invalid username or password")
		return
	}

	accessTokenCookie := http.Cookie{
		Name:     handlers.cfg.Jwt.Cookie.Name,
		Value:    userTokenDTO.Token,
		MaxAge:   handlers.cfg.Jwt.Cookie.MaxAge,
		Secure:   handlers.cfg.Jwt.Cookie.Secure,
		HttpOnly: handlers.cfg.Jwt.Cookie.HttpOnly,
	}

	http.SetCookie(response, &accessTokenCookie)

	response.Header().Set("Content-Type", "application/json")
	if err := utils.WriteResponse(response, http.StatusOK, userTokenDTO); err != nil {
		handlers.logger.Errorf("error encoding userTokenDTO: %v", err)
		utils.JSONError(response, http.StatusInternalServerError, "failed to return token")
		return
	}
}

// Login godoc
// @Summary Login user
// @Description Authenticate user with username and password
// @Accept json
// @Produce json
// @Param user body models.User true "User login details"
// @Success 200 {object} dto.UserTokenDTO "Login successful"
// @Failure 400 {object} utils.ErrorResponse "Invalid request body or missing fields"
// @Failure 401 {object} utils.ErrorResponse "Invalid username or password"
// @Failure 500 {object} utils.ErrorResponse "Failed to return token"
// @Router /api/v1/auth/login [post]
func (handlers *userHandlers) Login(response http.ResponseWriter, request *http.Request) {
	var user models.User

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		utils.JSONError(response, http.StatusBadRequest, "invalid request body")
		return
	}

	if user.Username == "" || user.Password == "" {
		utils.JSONError(response, http.StatusBadRequest, "username and password are required")
		return
	}

	userTokenDTO, err := handlers.usecase.Login(request.Context(), &user)
	if err != nil {
		handlers.logger.Warnf("failed to login user: %v", err)
		utils.JSONError(response, http.StatusUnauthorized, "invalid username or password")
		return
	}

	accessTokenCookie := http.Cookie{
		Name:     handlers.cfg.Jwt.Cookie.Name,
		Value:    userTokenDTO.Token,
		MaxAge:   handlers.cfg.Jwt.Cookie.MaxAge,
		Secure:   handlers.cfg.Jwt.Cookie.Secure,
		HttpOnly: handlers.cfg.Jwt.Cookie.HttpOnly,
	}

	http.SetCookie(response, &accessTokenCookie)

	response.Header().Set("Content-Type", "application/json")
	if err := utils.WriteResponse(response, http.StatusOK, userTokenDTO); err != nil {
		handlers.logger.Errorf("failed to encode userTokenDTO: %v", err)
		utils.JSONError(response, http.StatusInternalServerError, "failed to return token")
		return
	}
}

// Logout godoc
// @Summary Log out user
// @Description Log out user and clear access token cookie
// @Success 200 {object} utils.MessageResponse "Logout successful"
// @Failure 500 {object} utils.ErrorResponse "Failed to log out"
// @Router /api/v1/auth/logout [post]
func (handlers *userHandlers) Logout(response http.ResponseWriter, request *http.Request) {
	accessTokenCookie := http.Cookie{
		Name:     handlers.cfg.Jwt.Cookie.Name,
		Value:    "",
		MaxAge:   -1,
		Secure:   handlers.cfg.Jwt.Cookie.Secure,
		HttpOnly: handlers.cfg.Jwt.Cookie.HttpOnly,
	}

	http.SetCookie(response, &accessTokenCookie)

	response.Header().Set("Content-Type", "application/json")
	if err := utils.WriteResponse(response, http.StatusOK, "successfully logged out"); err != nil {
		handlers.logger.Errorf("error encoding logout response: %v", err)
		utils.JSONError(response, http.StatusInternalServerError, "failed to log out")
		return
	}
}

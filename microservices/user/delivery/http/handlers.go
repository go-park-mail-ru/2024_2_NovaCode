package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/user"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/user/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/content"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/csrf"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3"
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
	requestID := request.Context().Value(utils.RequestIDKey{})
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(response).Encode(utils.NewMessageResponse("OK")); err != nil {
		handlers.logger.Error(fmt.Sprintf("error encoding health response: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "failed to get health status")
		return
	}
}

// Register godoc
// @Tags Authentication
// @Summary Register a new user
// @Description Register a new user with a unique username, email, and password. On success, returns a user token.
// @Accept json
// @Produce json
// @Success 200 {object} dto.UserTokenDTO "User registration successful with token"
// @Failure 400 {object} utils.ErrorResponse "Invalid request body or missing fields"
// @Failure 500 {object} utils.ErrorResponse "Failed to return token"
// @Router /api/v1/auth/register [post]
func (handlers *userHandlers) Register(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	var regDTO dto.RegisterDTO

	if err := json.NewDecoder(request.Body).Decode(&regDTO); err != nil {
		utils.JSONError(response, http.StatusBadRequest, "invalid request body")
		handlers.logger.Errorf("error encoding register response: %v", err)
		return
	}

	if regDTO.Role == "" {
		regDTO.Role = "regular"
	}

	if err := regDTO.Validate(); err != nil {
		handlers.logger.Warnf(fmt.Sprintf("validation error: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("validation error: %v", err))
		return
	}

	user := dto.NewUserFromRegisterDTO(&regDTO)

	err := user.Sanitize()
	if err != nil {
		handlers.logger.Warn(fmt.Sprintf("sanitized: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("invalid username or password: %v", err))
		return
	}

	userTokenDTO, err := handlers.usecase.Register(request.Context(), user)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("failed to register user: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, "invalid username or password")
		return
	}

	accessTokenCookie := http.Cookie{
		Path:     "/",
		Name:     handlers.cfg.Jwt.Cookie.Name,
		Value:    userTokenDTO.Token,
		MaxAge:   handlers.cfg.Jwt.Cookie.MaxAge,
		Secure:   handlers.cfg.Jwt.Cookie.Secure,
		HttpOnly: handlers.cfg.Jwt.Cookie.HttpOnly,
	}

	http.SetCookie(response, &accessTokenCookie)

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(userTokenDTO); err != nil {
		handlers.logger.Errorf("error encoding userTokenDTO: %v", err)
		utils.JSONError(response, http.StatusInternalServerError, "failed to return token")
		return
	}

	response.WriteHeader(http.StatusOK)
}

// Login godoc
// @Tags Authentication
// @Summary User Login
// @Description Authenticate a user using their username and password. On success, returns an authentication token.
// @Accept json
// @Produce json
// @Success 200 {object} dto.UserTokenDTO "Login successful with token"
// @Failure 400 {object} utils.ErrorResponse "Invalid request body or missing fields"
// @Failure 401 {object} utils.ErrorResponse "Invalid username or password"
// @Failure 500 {object} utils.ErrorResponse "Failed to return token"
// @Router /api/v1/auth/login [post]
func (handlers *userHandlers) Login(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	var loginDTO dto.LoginDTO

	if err := json.NewDecoder(request.Body).Decode(&loginDTO); err != nil {
		utils.JSONError(response, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := loginDTO.Validate(); err != nil {
		handlers.logger.Warnf(fmt.Sprintf("validation error: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("validation error: %v", err))
		return
	}

	user := dto.NewUserFromLoginDTO(&loginDTO)

	err := user.Sanitize()
	if err != nil {
		handlers.logger.Warn(fmt.Sprintf("sanitized user: %v", err), requestID)
		utils.JSONError(response, http.StatusUnauthorized, fmt.Sprintf("invalid username or password: %v", err))
	}

	userTokenDTO, err := handlers.usecase.Login(request.Context(), user)
	if err != nil {
		handlers.logger.Warn(fmt.Sprintf("failed to login user: %v", err), requestID)
		utils.JSONError(response, http.StatusUnauthorized, "invalid username or password")
		return
	}

	accessTokenCookie := http.Cookie{
		Path:     "/",
		Name:     handlers.cfg.Jwt.Cookie.Name,
		Value:    userTokenDTO.Token,
		MaxAge:   handlers.cfg.Jwt.Cookie.MaxAge,
		Secure:   handlers.cfg.Jwt.Cookie.Secure,
		HttpOnly: handlers.cfg.Jwt.Cookie.HttpOnly,
	}

	http.SetCookie(response, &accessTokenCookie)

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(userTokenDTO); err != nil {
		handlers.logger.Error(fmt.Sprintf("failed to encode userTokenDTO: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "failed to return token")
		return
	}

	response.WriteHeader(http.StatusOK)
}

// Logout godoc
// @Tags Authentication
// @Summary Log out user
// @Description Clears the access token cookie to log the user out.
// @Success 200 {object} utils.MessageResponse "Logout successful"
// @Failure 500 {object} utils.ErrorResponse "Failed to log out"
// @Router /api/v1/auth/logout [post]
func (handlers *userHandlers) Logout(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	accessTokenCookie := http.Cookie{
		Path:     "/",
		Name:     handlers.cfg.Jwt.Cookie.Name,
		Value:    "",
		MaxAge:   -1,
		Secure:   handlers.cfg.Jwt.Cookie.Secure,
		HttpOnly: handlers.cfg.Jwt.Cookie.HttpOnly,
	}

	http.SetCookie(response, &accessTokenCookie)

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(utils.NewMessageResponse("successfully logged out")); err != nil {
		handlers.logger.Error(fmt.Sprintf("error encoding logout response: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "failed to log out")
		return
	}

	response.WriteHeader(http.StatusOK)
}

// GetCSRFToken godoc
// @Tags Authentication
// @Summary Generate a CSRF token
// @Description Generates a CSRF token for the authenticated user
// @Success 200 {object} utils.MessageResponse "CSRF token generated successfully"
// @Failure 401 {object} utils.ErrorResponse "unauthorized"
// @Failure 403 {object} utils.ErrorResponse "forbidden"
// @Router /api/v1/auth/csrf [get]
func (handlers *userHandlers) GetCSRFToken(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("user id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "user id not found")
		return
	}

	token := csrf.Generate(userID.String(), handlers.cfg.CSRF.Salt)

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(utils.NewCSRFResponse(token)); err != nil {
		handlers.logger.Error(fmt.Sprintf("error encoding csrf token response: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "failed to log out")
		return
	}

	response.WriteHeader(http.StatusOK)
}

// Update godoc
// @Tags User
// @Summary Update user details
// @Description Update user profile information such as username and email. Requires a valid user ID in the request context.
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} dto.UserDTO "User update successful"
// @Failure 400 {object} utils.ErrorResponse "Invalid request body or missing fields"
// @Failure 401 {object} utils.ErrorResponse "User not authenticated"
// @Failure 403 {object} utils.ErrorResponse "Not enough permissions to update user details"
// @Failure 404 {object} utils.ErrorResponse "User ID not found in context"
// @Failure 500 {object} utils.ErrorResponse "Failed to update user details"
// @Router /api/v1/users/{user_id} [put]
func (handlers *userHandlers) Update(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	var updateDTO dto.UpdateDTO

	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("user id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "user id not found")
		return
	}
	updateDTO.UserID = userID

	if err := json.NewDecoder(request.Body).Decode(&updateDTO); err != nil {
		utils.JSONError(response, http.StatusBadRequest, "invalid request body")
		return
	}

	user := dto.NewUserFromUpdateDTO(&updateDTO)
	updatedUserDTO, err := handlers.usecase.Update(request.Context(), user)
	if err != nil {
		handlers.logger.Warn(fmt.Sprintf("failed to update user: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "failed to update user details")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(updatedUserDTO); err != nil {
		handlers.logger.Error(fmt.Sprintf("error encoding updated user response: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "failed to return updated user details")
		return
	}

	response.WriteHeader(http.StatusOK)
}

// UploadImage godoc
// @Tags User
// @Summary Upload profile imag
// @Description Upload a profile image for the user. The image file should be in a supported image format.
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Profile image file"
// @Param user_id path string true "User ID"
// @Success 200 {object} dto.UserDTO "Image uploaded successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid file format or missing file"
// @Failure 401 {object} utils.ErrorResponse "User not authenticated"
// @Failure 403 {object} utils.ErrorResponse "Not enough permissions to upload image"
// @Failure 500 {object} utils.ErrorResponse "Failed to upload image"
// @Router /api/v1/users/{user_id}/image [post]
func (handlers *userHandlers) UploadImage(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("user id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "failed to update user details")
		return
	}

	file, header, err := request.FormFile("file")
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, "failed to get file from request")
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("failed to read file: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "failed to read file")
		return
	}

	contentType, err := content.IsImage(fileBytes)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("invalid content type: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, "invalid content type")
		return
	}

	upload := s3.Upload{
		Bucket:      "avatars",
		File:        bytes.NewReader(fileBytes),
		Filename:    header.Filename,
		Size:        header.Size,
		ContentType: contentType,
	}

	userDTO, err := handlers.usecase.UploadImage(request.Context(), userID, upload)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("failed to upload image: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "failed to upload image")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(userDTO); err != nil {
		handlers.logger.Error(fmt.Sprintf("error encoding updated user response: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "failed to encode response")
		return
	}

	response.WriteHeader(http.StatusOK)
}

// GetUserByUsername godoc
// @Tags User
// @Summary Get user by username
// @Description Retrieves public profile details for the specified username.
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} dto.PublicUserDTO "User details retrieved successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid or missing username"
// @Failure 404 {object} utils.ErrorResponse "User not found"
// @Router /api/v1/users/{username} [get]
func (handlers *userHandlers) GetUserByUsername(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	username, ok := vars["username"]
	if !ok {
		utils.JSONError(response, http.StatusBadRequest, "username is required")
		return
	}

	userDTO, err := handlers.usecase.GetByUsername(request.Context(), username)
	if err != nil {
		utils.JSONError(response, http.StatusNotFound, "username not found")
		return
	}
	publicUserDTO := dto.NewPublicUserDTO(userDTO)

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(publicUserDTO); err != nil {
		utils.JSONError(response, http.StatusInternalServerError, "failed to encode response")
		return
	}

	response.WriteHeader(http.StatusOK)
}

// GetMe godoc
// @Tags User
// @Summary Get current user details
// @Description Retrieves profile details for the currently authenticated user.
// @Produce json
// @Success 200 {object} dto.UserDTO "User details retrieved successfully"
// @Failure 401 {object} utils.ErrorResponse "User not authenticated"
// @Failure 500 {object} utils.ErrorResponse "Failed to retrieve user details"
// @Router /api/v1/users/me [get]
func (handlers *userHandlers) GetMe(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("user id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "user id not found")
		return
	}

	userDTO, err := handlers.usecase.GetByID(request.Context(), userID)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, "user id not found")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(userDTO); err != nil {
		utils.JSONError(response, http.StatusInternalServerError, "failed to encode response")
		return
	}

	response.WriteHeader(http.StatusOK)
}

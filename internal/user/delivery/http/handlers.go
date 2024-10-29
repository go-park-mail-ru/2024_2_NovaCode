package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/user"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/user/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/content"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

type contextKey string

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
	response.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(response).Encode(utils.NewMessageResponse("OK")); err != nil {
		handlers.logger.Errorf("error encoding health response: %v", err)
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
// @Param user body models.User true "User registration details"
// @Success 200 {object} dto.UserTokenDTO "User registration successful with token"
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
// @Param user body models.User true "User login details" example({ "username": "john_doe", "password": "password123" })
// @Success 200 {object} dto.UserTokenDTO "Login successful with token"
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
	if err := json.NewEncoder(response).Encode(userTokenDTO); err != nil {
		handlers.logger.Errorf("failed to encode userTokenDTO: %v", err)
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
	accessTokenCookie := http.Cookie{
		Name:     handlers.cfg.Jwt.Cookie.Name,
		Value:    "",
		MaxAge:   -1,
		Secure:   handlers.cfg.Jwt.Cookie.Secure,
		HttpOnly: handlers.cfg.Jwt.Cookie.HttpOnly,
	}

	http.SetCookie(response, &accessTokenCookie)

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(utils.NewMessageResponse("successfully logged out")); err != nil {
		handlers.logger.Errorf("error encoding logout response: %v", err)
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
// @Param user body models.User true "Updated user details" example({ "username": "new_username", "email": "new_email@example.com" })
// @Success 200 {object} dto.UserDTO "User update successful"
// @Failure 400 {object} utils.ErrorResponse "Invalid request body or missing fields"
// @Failure 401 {object} utils.ErrorResponse "User not authenticated"
// @Failure 403 {object} utils.ErrorResponse "Not enough permissions to update user details"
// @Failure 404 {object} utils.ErrorResponse "User ID not found in context"
// @Failure 500 {object} utils.ErrorResponse "Failed to update user details"
// @Router /api/v1/users/{user_id}/update [put]
func (handlers *userHandlers) Update(response http.ResponseWriter, request *http.Request) {
	var user models.User

	userID, ok := request.Context().Value(contextKey("userID")).(uuid.UUID)
	if !ok {
		handlers.logger.Errorf("user id not found in context")
		utils.JSONError(response, http.StatusNotFound, "user id not found")
		return
	}
	user.UserID = userID

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		utils.JSONError(response, http.StatusBadRequest, "invalid request body")
		return
	}

	updatedUserDTO, err := handlers.usecase.Update(request.Context(), &user)
	if err != nil {
		handlers.logger.Warnf("failed to update user: %v", err)
		utils.JSONError(response, http.StatusInternalServerError, "failed to update user details")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(updatedUserDTO); err != nil {
		handlers.logger.Errorf("error encoding updated user response: %v", err)
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
	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Errorf("user id not found in context")
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
		handlers.logger.Errorf("failed to read file: %v", err)
		utils.JSONError(response, http.StatusInternalServerError, "failed to read file")
		return
	}

	contentType, err := content.IsImage(fileBytes)
	if err != nil {
		handlers.logger.Errorf("invalid content type: %v", err)
		utils.JSONError(response, http.StatusBadRequest, "invalid content type")
		return
	}

	upload := s3.Upload{
		Bucket:      "users",
		File:        bytes.NewReader(fileBytes),
		Filename:    header.Filename,
		Size:        header.Size,
		ContentType: contentType,
	}

	userDTO, err := handlers.usecase.UploadImage(request.Context(), userID, upload)
	if err != nil {
		handlers.logger.Errorf("failed to upload image: %v", err)
		utils.JSONError(response, http.StatusInternalServerError, "failed to upload image")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(userDTO); err != nil {
		handlers.logger.Errorf("error encoding updated user response: %v", err)
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
	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Errorf("user id not found in context")
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

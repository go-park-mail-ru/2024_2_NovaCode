package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/user/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/user/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/csrf"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestUserHandlers_Health(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}

	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mock.NewMockUsecase(ctrl)
	userHandlers := NewUserHandlers(&cfg.Service.Auth, usecaseMock, logger)

	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	response := httptest.NewRecorder()

	userHandlers.Health(response, request)

	result := response.Result()

	assert.Equal(t, http.StatusOK, result.StatusCode)

	expectedBody := `{"message": "OK"}`
	assert.JSONEq(t, expectedBody, response.Body.String())
}

func TestUserHandlers_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Auth: config.AuthConfig{
				Jwt: config.JwtConfig{
					Cookie: config.JwtCookieConfig{
						Name:     "access_token",
						MaxAge:   3600,
						Secure:   true,
						HttpOnly: true,
					},
				},
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mock.NewMockUsecase(ctrl)
	userHandlers := NewUserHandlers(&cfg.Service.Auth, usecaseMock, logger)

	t.Run("successful registration", func(t *testing.T) {
		user := models.User{
			Username: "test_user",
			Email:    "test@example.com",
			Password: "password",
			Role:     "regular",
		}

		userTokenDTO := &dto.UserTokenDTO{
			Token: "test_token",
		}

		usecaseMock.EXPECT().Register(gomock.Any(), &user).Return(userTokenDTO, nil)

		body, _ := json.Marshal(user)
		request := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		response := httptest.NewRecorder()

		userHandlers.Register(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		cookie := response.Result().Cookies()[0]
		assert.Equal(t, cfg.Service.Auth.Jwt.Cookie.Name, cookie.Name)
		assert.Equal(t, userTokenDTO.Token, cookie.Value)
		assert.Equal(t, cfg.Service.Auth.Jwt.Cookie.MaxAge, cookie.MaxAge)
		assert.True(t, cookie.HttpOnly)
		assert.True(t, cookie.Secure)
	})

	t.Run("invalid request body", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString("invalid json"))
		response := httptest.NewRecorder()

		userHandlers.Register(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
	})

	t.Run("missing required fields", func(t *testing.T) {
		user := models.User{
			Username: "",
			Email:    "test@example.com",
			Password: "password",
		}

		body, _ := json.Marshal(user)
		request := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		response := httptest.NewRecorder()

		userHandlers.Register(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
	})

	t.Run("usecase registration error", func(t *testing.T) {
		user := models.User{
			Username: "test_user",
			Email:    "test@example.com",
			Password: "password",
			Role:     "regular",
		}

		usecaseMock.EXPECT().Register(gomock.Any(), &user).Return(nil, errors.New("registration error"))

		body, _ := json.Marshal(user)
		request := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		response := httptest.NewRecorder()

		userHandlers.Register(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
	})
}

func TestUserHandlers_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Auth: config.AuthConfig{
				Jwt: config.JwtConfig{
					Cookie: config.JwtCookieConfig{
						Name:     "access_token",
						MaxAge:   3600,
						Secure:   true,
						HttpOnly: true,
					},
				},
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mock.NewMockUsecase(ctrl)
	userHandlers := NewUserHandlers(&cfg.Service.Auth, usecaseMock, logger)

	t.Run("successful login", func(t *testing.T) {
		user := models.User{
			Username: "test_user",
			Password: "password",
		}

		userTokenDTO := &dto.UserTokenDTO{
			Token: "test_token",
		}

		usecaseMock.EXPECT().Login(gomock.Any(), &user).Return(userTokenDTO, nil)

		body, _ := json.Marshal(user)
		request := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		response := httptest.NewRecorder()

		userHandlers.Login(response, request)

		result := response.Result()
		assert.Equal(t, http.StatusOK, result.StatusCode)

		cookie := response.Result().Cookies()[0]
		assert.Equal(t, cfg.Service.Auth.Jwt.Cookie.Name, cookie.Name)
		assert.Equal(t, userTokenDTO.Token, cookie.Value)
		assert.Equal(t, cfg.Service.Auth.Jwt.Cookie.MaxAge, cookie.MaxAge)
		assert.True(t, cookie.HttpOnly)
		assert.True(t, cookie.Secure)

		var responseDTO dto.UserTokenDTO
		err := json.NewDecoder(response.Body).Decode(&responseDTO)
		assert.NoError(t, err)
		assert.Equal(t, userTokenDTO.Token, responseDTO.Token)
	})

	t.Run("invalid request body", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString("invalid json"))
		response := httptest.NewRecorder()

		userHandlers.Login(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
	})

	t.Run("missing required fields", func(t *testing.T) {
		user := models.User{
			Username: "",
			Password: "",
		}

		body, _ := json.Marshal(user)
		request := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		response := httptest.NewRecorder()

		userHandlers.Login(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
	})

	t.Run("usecase login error", func(t *testing.T) {
		user := models.User{
			Username: "test_user",
			Password: "password",
		}

		usecaseMock.EXPECT().Login(gomock.Any(), &user).Return(nil, errors.New("login error"))

		body, _ := json.Marshal(user)
		request := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		response := httptest.NewRecorder()

		userHandlers.Login(response, request)

		assert.Equal(t, http.StatusUnauthorized, response.Result().StatusCode)
	})
}

func TestUserHandlers_Logout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Auth: config.AuthConfig{
				Jwt: config.JwtConfig{
					Cookie: config.JwtCookieConfig{
						Name:     "access_token",
						MaxAge:   3600,
						Secure:   true,
						HttpOnly: true,
					},
				},
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	handlers := NewUserHandlers(&cfg.Service.Auth, nil, logger)

	t.Run("successful logout", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/logout", bytes.NewBuffer(nil))
		response := httptest.NewRecorder()

		handlers.Logout(response, request)

		result := response.Result()
		assert.Equal(t, http.StatusOK, result.StatusCode)

		cookie := response.Result().Cookies()[0]
		assert.Equal(t, cfg.Service.Auth.Jwt.Cookie.Name, cookie.Name)
		assert.Equal(t, "", cookie.Value)
		assert.Equal(t, -1, cookie.MaxAge)

		var responseMsg map[string]string
		err := json.NewDecoder(response.Body).Decode(&responseMsg)
		assert.NoError(t, err)
		assert.Equal(t, "successfully logged out", responseMsg["message"])
	})
}

func TestUserHandlers_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Auth: config.AuthConfig{
				Jwt: config.JwtConfig{
					Cookie: config.JwtCookieConfig{
						Name:     "access_token",
						MaxAge:   3600,
						Secure:   true,
						HttpOnly: true,
					},
				},
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mock.NewMockUsecase(ctrl)
	userHandlers := NewUserHandlers(&cfg.Service.Auth, usecaseMock, logger)

	userID := uuid.New()

	t.Run("successful update", func(t *testing.T) {
		user := models.User{
			UserID:   userID,
			Email:    "updated@example.com",
			Password: "newpassword",
		}

		userDTO := &dto.UserDTO{
			ID:       user.UserID,
			Username: user.Username,
		}

		usecaseMock.EXPECT().Update(gomock.Any(), gomock.Eq(&user)).Return(userDTO, nil)

		body, _ := json.Marshal(user)
		request := httptest.NewRequest(http.MethodPut, "/update", bytes.NewBuffer(body))
		ctx := context.WithValue(request.Context(), utils.UserIDKey{}, userID)
		request = request.WithContext(ctx)

		response := httptest.NewRecorder()

		userHandlers.Update(response, request)

		assert.Equal(t, http.StatusOK, response.Result().StatusCode)

		var responseMsg dto.UserDTO
		err := json.NewDecoder(response.Body).Decode(&responseMsg)
		assert.NoError(t, err)
		assert.Equal(t, userDTO, &responseMsg)
	})

	t.Run("invalid request body", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPut, "/update", bytes.NewBufferString("invalid json"))
		ctx := context.WithValue(request.Context(), utils.UserIDKey{}, userID)
		request = request.WithContext(ctx)
		response := httptest.NewRecorder()

		userHandlers.Update(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
	})

	t.Run("usecase update error", func(t *testing.T) {
		user := models.User{
			UserID:   userID,
			Username: "updated_user",
			Email:    "updated@example.com",
			Password: "newpassword",
		}

		usecaseMock.EXPECT().Update(gomock.Any(), gomock.Eq(&user)).Return(nil, errors.New("update error"))

		body, _ := json.Marshal(user)
		request := httptest.NewRequest(http.MethodPut, "/update", bytes.NewBuffer(body))
		ctx := context.WithValue(request.Context(), utils.UserIDKey{}, userID)
		request = request.WithContext(ctx)

		response := httptest.NewRecorder()

		userHandlers.Update(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Result().StatusCode)
	})
}

func TestUserHandlers_GetUserByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Auth: config.AuthConfig{
				Jwt: config.JwtConfig{
					Cookie: config.JwtCookieConfig{
						Name:     "access_token",
						MaxAge:   3600,
						Secure:   true,
						HttpOnly: true,
					},
				},
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mock.NewMockUsecase(ctrl)
	userHandlers := NewUserHandlers(nil, usecaseMock, logger)

	t.Run("successful get user by username", func(t *testing.T) {
		username := "test_user"
		userDTO := &dto.UserDTO{
			Username: username,
			Email:    "test@example.com",
		}
		publicUserDTO := dto.NewPublicUserDTO(userDTO)

		usecaseMock.EXPECT().GetByUsername(gomock.Any(), username).Return(userDTO, nil)

		request := httptest.NewRequest(http.MethodGet, "/users/"+username, nil)
		request = mux.SetURLVars(request, map[string]string{"username": username})
		response := httptest.NewRecorder()

		userHandlers.GetUserByUsername(response, request)

		result := response.Result()
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, "application/json", result.Header.Get("Content-Type"))

		var responseDTO dto.PublicUserDTO
		err := json.NewDecoder(response.Body).Decode(&responseDTO)
		assert.NoError(t, err)
		assert.Equal(t, *publicUserDTO, responseDTO)
	})

	t.Run("missing username", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/users/", nil)
		response := httptest.NewRecorder()

		userHandlers.GetUserByUsername(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
	})

	t.Run("user not found", func(t *testing.T) {
		username := "nonexistent_user"
		usecaseMock.EXPECT().GetByUsername(gomock.Any(), username).Return(nil, errors.New("user not found"))

		request := httptest.NewRequest(http.MethodGet, "/users/"+username, nil)
		request = mux.SetURLVars(request, map[string]string{"username": username})
		response := httptest.NewRecorder()

		userHandlers.GetUserByUsername(response, request)

		assert.Equal(t, http.StatusNotFound, response.Result().StatusCode)
	})
}

func TestUserHandlers_GetMe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Auth: config.AuthConfig{
				Jwt: config.JwtConfig{
					Cookie: config.JwtCookieConfig{
						Name:     "access_token",
						MaxAge:   3600,
						Secure:   true,
						HttpOnly: true,
					},
				},
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mock.NewMockUsecase(ctrl)
	userHandlers := NewUserHandlers(nil, usecaseMock, logger)

	t.Run("successful get me", func(t *testing.T) {
		userID := uuid.New()
		userDTO := &dto.UserDTO{
			ID:       userID,
			Username: "test_user",
			Email:    "test@example.com",
		}

		usecaseMock.EXPECT().GetByID(gomock.Any(), userID).Return(userDTO, nil)

		request := httptest.NewRequest(http.MethodGet, "/me", nil)
		ctx := context.WithValue(request.Context(), utils.UserIDKey{}, userID)
		request = request.WithContext(ctx)

		response := httptest.NewRecorder()

		userHandlers.GetMe(response, request)

		result := response.Result()
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, "application/json", result.Header.Get("Content-Type"))

		var responseDTO dto.UserDTO
		err := json.NewDecoder(response.Body).Decode(&responseDTO)
		assert.NoError(t, err)
		assert.Equal(t, *userDTO, responseDTO)
	})

	t.Run("user id not found in context", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/me", nil)
		response := httptest.NewRecorder()

		userHandlers.GetMe(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
	})

	t.Run("user id not found in database", func(t *testing.T) {
		userID := uuid.New()

		usecaseMock.EXPECT().GetByID(gomock.Any(), userID).Return(nil, errors.New("user not found"))

		request := httptest.NewRequest(http.MethodGet, "/me", nil)
		ctx := context.WithValue(request.Context(), utils.UserIDKey{}, userID)
		request = request.WithContext(ctx)
		response := httptest.NewRecorder()

		userHandlers.GetMe(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
	})
}

func TestUserHandlers_GetCSRFToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "debug",
				Format: "json",
			},
			Auth: config.AuthConfig{
				CSRF: config.CSRFConfig{
					HeaderName: "X-CSRF-Token",
					Salt:       "yD5pwD0JG03NxFAz9VAOtbba6I7kPv2deg3C7SpEZUk=",
				},
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mock.NewMockUsecase(ctrl)
	userHandlers := NewUserHandlers(&cfg.Service.Auth, usecaseMock, logger)

	t.Run("successful get csrf token", func(t *testing.T) {
		userID := uuid.New()
		expectedToken := csrf.Generate(userID.String(), cfg.Service.Auth.CSRF.Salt)

		request := httptest.NewRequest(http.MethodGet, "/csrf", nil)
		ctx := context.WithValue(request.Context(), utils.UserIDKey{}, userID)
		request = request.WithContext(ctx)

		response := httptest.NewRecorder()

		userHandlers.GetCSRFToken(response, request)

		result := response.Result()
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, expectedToken, result.Header.Get(cfg.Service.Auth.CSRF.HeaderName))
	})

	t.Run("user id not found in context", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/csrf", nil)
		response := httptest.NewRecorder()

		userHandlers.GetCSRFToken(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
	})
}

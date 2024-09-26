package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/user/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/user/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

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

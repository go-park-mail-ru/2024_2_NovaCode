package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/daronenko/auth/config"
	"github.com/daronenko/auth/internal/models"
	"github.com/daronenko/auth/internal/user/mock"
	"github.com/daronenko/auth/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsecase_Register(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Auth: config.AuthConfig{
			Jwt: config.JwtConfig{
				Secret: "secret",
			},
		},
		Logger: config.LoggerConfig{
			Level:  "info",
			Format: "json",
		},
	}

	logger := logger.New(cfg)
	repoMock := mock.NewMockRepo(ctrl)
	userUsecase := NewUserUsecase(cfg, repoMock, logger)

	user := &models.User{
		Username: "test_user",
		Email:    "test@email.com",
		Password: "password",
	}

	ctx := context.Background()

	repoMock.EXPECT().FindByUsername(ctx, gomock.Eq(user.Username)).Return(nil, sql.ErrNoRows)
	repoMock.EXPECT().FindByEmail(ctx, gomock.Eq(user.Email)).Return(nil, sql.ErrNoRows)
	repoMock.EXPECT().Insert(ctx, gomock.Eq(user)).Return(user, nil)

	userToken, err := userUsecase.Register(context.Background(), user)
	require.NoError(t, err)
	require.NotNil(t, userToken)
	require.Nil(t, err)
}

func TestUsecase_Register_UsernameAlreadyExists(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Auth: config.AuthConfig{
			Jwt: config.JwtConfig{
				Secret: "secret",
			},
		},
		Logger: config.LoggerConfig{
			Level:  "info",
			Format: "json",
		},
	}

	logger := logger.New(cfg)
	repoMock := mock.NewMockRepo(ctrl)
	userUsecase := NewUserUsecase(cfg, repoMock, logger)

	user := &models.User{
		Username: "test_user",
		Email:    "test@email.com",
		Password: "password",
	}

	ctx := context.Background()

	existingUser := &models.User{
		Username: "test_user",
		Email:    "existing@email.com",
	}
	repoMock.EXPECT().FindByUsername(ctx, gomock.Eq(user.Username)).Return(existingUser, nil)

	userToken, err := userUsecase.Register(ctx, user)
	require.Error(t, err)
	require.Nil(t, userToken)
	require.EqualError(t, err, "user with that username already exists")
}

func TestUsecase_Register_EmailAlreadyExists(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Auth: config.AuthConfig{
			Jwt: config.JwtConfig{
				Secret: "secret",
			},
		},
		Logger: config.LoggerConfig{
			Level:  "info",
			Format: "json",
		},
	}

	logger := logger.New(cfg)
	repoMock := mock.NewMockRepo(ctrl)
	userUsecase := NewUserUsecase(cfg, repoMock, logger)

	user := &models.User{
		Username: "test_user",
		Email:    "test@email.com",
		Password: "password",
	}

	ctx := context.Background()

	repoMock.EXPECT().FindByUsername(ctx, gomock.Eq(user.Username)).Return(nil, sql.ErrNoRows)

	existingUser := &models.User{
		Username: "existing_user",
		Email:    "test@email.com",
	}
	repoMock.EXPECT().FindByEmail(ctx, gomock.Eq(user.Email)).Return(existingUser, nil)

	userToken, err := userUsecase.Register(ctx, user)
	require.Error(t, err)
	require.Nil(t, userToken)
	require.EqualError(t, err, "user with that email already exists")
}

func TestUsecase_Register_InsertError(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Auth: config.AuthConfig{
			Jwt: config.JwtConfig{
				Secret: "secret",
			},
		},
		Logger: config.LoggerConfig{
			Level:  "info",
			Format: "json",
		},
	}

	logger := logger.New(cfg)
	repoMock := mock.NewMockRepo(ctrl)
	userUsecase := NewUserUsecase(cfg, repoMock, logger)

	user := &models.User{
		Username: "test_user",
		Email:    "test@email.com",
		Password: "password",
	}

	ctx := context.Background()

	repoMock.EXPECT().FindByUsername(ctx, gomock.Eq(user.Username)).Return(nil, sql.ErrNoRows)
	repoMock.EXPECT().FindByEmail(ctx, gomock.Eq(user.Email)).Return(nil, sql.ErrNoRows)

	repoMock.EXPECT().Insert(ctx, gomock.Eq(user)).Return(nil, fmt.Errorf("insert error"))

	userToken, err := userUsecase.Register(ctx, user)
	require.Error(t, err)
	require.Nil(t, userToken)
	require.EqualError(t, err, "failed to create user: insert error")
}

func TestUsecase_Login(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Auth: config.AuthConfig{
			Jwt: config.JwtConfig{
				Secret: "secret",
			},
		},
		Logger: config.LoggerConfig{
			Level:  "info",
			Format: "json",
		},
	}

	logger := logger.New(cfg)
	repoMock := mock.NewMockRepo(ctrl)
	userUsecase := NewUserUsecase(cfg, repoMock, logger)

	password := "password"

	user := &models.User{
		Username: "test_user",
		Email:    "test@email.com",
		Password: password,
	}

	userMock := &models.User{
		Username: "test_user",
		Email:    "test@email.com",
		Password: password,
	}
	err := userMock.HashPassword()
	require.NoError(t, err)

	ctx := context.Background()

	repoMock.EXPECT().FindByUsername(ctx, gomock.Eq(user.Username)).Return(userMock, nil)

	userToken, err := userUsecase.Login(ctx, user)
	require.NoError(t, err)
	require.NotNil(t, userToken)
	require.Nil(t, err)
}

func TestUsecase_Login_InvalidPassword(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Auth: config.AuthConfig{
			Jwt: config.JwtConfig{
				Secret: "secret",
			},
		},
		Logger: config.LoggerConfig{
			Level:  "info",
			Format: "json",
		},
	}

	logger := logger.New(cfg)
	repoMock := mock.NewMockRepo(ctrl)
	userUsecase := NewUserUsecase(cfg, repoMock, logger)

	password := "password"
	wrongPassword := "wrong_password"

	user := &models.User{
		Username: "test_user",
		Password: wrongPassword,
	}

	userMock := &models.User{
		Username: "test_user",
		Password: password,
	}
	err := userMock.HashPassword()
	require.NoError(t, err)

	ctx := context.Background()

	repoMock.EXPECT().FindByUsername(ctx, gomock.Eq(user.Username)).Return(userMock, nil)

	userToken, err := userUsecase.Login(ctx, user)
	require.Error(t, err)
	require.Nil(t, userToken)
	require.EqualError(t, err, "invalid username or password")
}

func TestUsecase_Login_UserNotFound(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Auth: config.AuthConfig{
			Jwt: config.JwtConfig{
				Secret: "secret",
			},
		},
		Logger: config.LoggerConfig{
			Level:  "info",
			Format: "json",
		},
	}

	logger := logger.New(cfg)
	repoMock := mock.NewMockRepo(ctrl)
	userUsecase := NewUserUsecase(cfg, repoMock, logger)

	user := &models.User{
		Username: "test_user",
		Password: "password",
	}

	ctx := context.Background()

	repoMock.EXPECT().FindByUsername(ctx, gomock.Eq(user.Username)).Return(nil, sql.ErrNoRows)

	userToken, err := userUsecase.Login(ctx, user)
	require.Error(t, err)
	require.Nil(t, userToken)
	require.EqualError(t, err, "invalid username or password")
}

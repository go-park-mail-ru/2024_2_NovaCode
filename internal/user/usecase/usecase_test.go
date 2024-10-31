package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/user/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUsecase_Register(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Auth: config.AuthConfig{
				Jwt: config.JwtConfig{
					Secret: "secret",
				},
			},
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	pgRepoMock := mock.NewMockPostgresRepo(ctrl)
	s3RepoMock := mock.NewMockS3Repo(ctrl)
	userUsecase := NewUserUsecase(&cfg.Service.Auth, &cfg.Minio, pgRepoMock, s3RepoMock, logger)

	user := &models.User{
		Username: "test_user",
		Email:    "test@email.com",
		Password: "password",
	}

	ctx := context.Background()

	pgRepoMock.EXPECT().FindByUsername(ctx, gomock.Eq(user.Username)).Return(nil, sql.ErrNoRows)
	pgRepoMock.EXPECT().FindByEmail(ctx, gomock.Eq(user.Email)).Return(nil, sql.ErrNoRows)
	pgRepoMock.EXPECT().Insert(ctx, gomock.Eq(user)).Return(user, nil)

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
		Service: config.ServiceConfig{
			Auth: config.AuthConfig{
				Jwt: config.JwtConfig{
					Secret: "secret",
				},
			},
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	pgRepoMock := mock.NewMockPostgresRepo(ctrl)
	s3RepoMock := mock.NewMockS3Repo(ctrl)
	userUsecase := NewUserUsecase(&cfg.Service.Auth, &cfg.Minio, pgRepoMock, s3RepoMock, logger)

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
	pgRepoMock.EXPECT().FindByUsername(ctx, gomock.Eq(user.Username)).Return(existingUser, nil)

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
		Service: config.ServiceConfig{
			Auth: config.AuthConfig{
				Jwt: config.JwtConfig{
					Secret: "secret",
				},
			},
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	pgRepoMock := mock.NewMockPostgresRepo(ctrl)
	s3RepoMock := mock.NewMockS3Repo(ctrl)
	userUsecase := NewUserUsecase(&cfg.Service.Auth, &cfg.Minio, pgRepoMock, s3RepoMock, logger)

	user := &models.User{
		Username: "test_user",
		Email:    "test@email.com",
		Password: "password",
	}

	ctx := context.Background()

	pgRepoMock.EXPECT().FindByUsername(ctx, gomock.Eq(user.Username)).Return(nil, sql.ErrNoRows)

	existingUser := &models.User{
		Username: "existing_user",
		Email:    "test@email.com",
	}
	pgRepoMock.EXPECT().FindByEmail(ctx, gomock.Eq(user.Email)).Return(existingUser, nil)

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
		Service: config.ServiceConfig{
			Auth: config.AuthConfig{
				Jwt: config.JwtConfig{
					Secret: "secret",
				},
			},
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	pgRepoMock := mock.NewMockPostgresRepo(ctrl)
	s3RepoMock := mock.NewMockS3Repo(ctrl)
	userUsecase := NewUserUsecase(&cfg.Service.Auth, &cfg.Minio, pgRepoMock, s3RepoMock, logger)

	user := &models.User{
		Username: "test_user",
		Email:    "test@email.com",
		Password: "password",
	}

	ctx := context.Background()

	pgRepoMock.EXPECT().FindByUsername(ctx, gomock.Eq(user.Username)).Return(nil, sql.ErrNoRows)
	pgRepoMock.EXPECT().FindByEmail(ctx, gomock.Eq(user.Email)).Return(nil, sql.ErrNoRows)

	pgRepoMock.EXPECT().Insert(ctx, gomock.Eq(user)).Return(nil, fmt.Errorf("insert error"))

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
		Service: config.ServiceConfig{
			Auth: config.AuthConfig{
				Jwt: config.JwtConfig{
					Secret: "secret",
				},
			},
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	pgRepoMock := mock.NewMockPostgresRepo(ctrl)
	s3RepoMock := mock.NewMockS3Repo(ctrl)
	userUsecase := NewUserUsecase(&cfg.Service.Auth, &cfg.Minio, pgRepoMock, s3RepoMock, logger)

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

	pgRepoMock.EXPECT().FindByUsername(ctx, gomock.Eq(user.Username)).Return(userMock, nil)

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
		Service: config.ServiceConfig{
			Auth: config.AuthConfig{
				Jwt: config.JwtConfig{
					Secret: "secret",
				},
			},
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	pgRepoMock := mock.NewMockPostgresRepo(ctrl)
	s3RepoMock := mock.NewMockS3Repo(ctrl)
	userUsecase := NewUserUsecase(&cfg.Service.Auth, &cfg.Minio, pgRepoMock, s3RepoMock, logger)

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

	pgRepoMock.EXPECT().FindByUsername(ctx, gomock.Eq(user.Username)).Return(userMock, nil)

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
		Service: config.ServiceConfig{
			Auth: config.AuthConfig{
				Jwt: config.JwtConfig{
					Secret: "secret",
				},
			},
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	pgRepoMock := mock.NewMockPostgresRepo(ctrl)
	s3RepoMock := mock.NewMockS3Repo(ctrl)
	userUsecase := NewUserUsecase(&cfg.Service.Auth, &cfg.Minio, pgRepoMock, s3RepoMock, logger)

	user := &models.User{
		Username: "test_user",
		Password: "password",
	}

	ctx := context.Background()

	pgRepoMock.EXPECT().FindByUsername(ctx, gomock.Eq(user.Username)).Return(nil, sql.ErrNoRows)

	userToken, err := userUsecase.Login(ctx, user)
	require.Error(t, err)
	require.Nil(t, userToken)
	require.EqualError(t, err, "invalid username or password")
}

func TestUsecase_Update(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.New(&config.LoggerConfig{Level: "info", Format: "json"})
	pgRepoMock := mock.NewMockPostgresRepo(ctrl)
	s3RepoMock := mock.NewMockS3Repo(ctrl)
	userUsecase := NewUserUsecase(nil, nil, pgRepoMock, s3RepoMock, logger)

	user := &models.User{
		UserID:   uuid.New(),
		Username: "updated_user",
		Email:    "updated_email@example.com",
		Password: "new_password",
	}

	ctx := context.Background()

	existingUser := &models.User{
		UserID:   user.UserID,
		Username: "old_user",
		Email:    "old_email@example.com",
		Password: "old_password",
	}
	pgRepoMock.EXPECT().FindByID(ctx, user.UserID).Return(existingUser, nil)

	pgRepoMock.EXPECT().Update(ctx, gomock.Any()).Return(user, nil)

	updatedUserDTO, err := userUsecase.Update(ctx, user)
	require.NoError(t, err)
	require.NotNil(t, updatedUserDTO)
	require.Equal(t, user.Username, updatedUserDTO.Username)
}

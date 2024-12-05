package postgres

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestInsert_Regular(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}
	logger := logger.New(&cfg.Service.Logger)

	postgresRepo := NewUserPostgresRepository(db, logger)

	columns := []string{"id", "username", "email", "password", "role", "image", "created_at", "updated_at"}

	userUUID := uuid.New()
	userMock := &models.User{
		UserID:   userUUID,
		Username: "test_user",
		Email:    "test@email.com",
		Password: "password",
		Role:     "regular",
		Image:    "avatar.jpeg",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userUUID,
		userMock.Username,
		userMock.Email,
		userMock.Password,
		userMock.Role,
		userMock.Image,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(insertUserQuery).WithArgs(
		userMock.Username,
		userMock.Email,
		userMock.Password,
		userMock.Role,
	).WillReturnRows(rows)

	insertedUser, err := postgresRepo.Insert(context.Background(), userMock)
	require.NotNil(t, insertedUser)
	require.NoError(t, err)
}

func TestInsert_Error(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}
	logger := logger.New(&cfg.Service.Logger)

	postgresRepo := NewUserPostgresRepository(db, logger)

	userMock := &models.User{
		Username: "test_user",
		Email:    "test@email.com",
		Password: "password",
		Role:     "regular",
		Image:    "avatar.jpeg",
	}

	mock.ExpectQuery(insertUserQuery).WithArgs(
		userMock.Username,
		userMock.Email,
		userMock.Password,
		userMock.Role,
		userMock.Image,
	).WillReturnError(fmt.Errorf("some error"))

	_, err = postgresRepo.Insert(context.Background(), userMock)
	require.Contains(t, err.Error(), "failed to insert user")
	require.Error(t, err)
}

func TestFindByID_Regular(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}
	logger := logger.New(&cfg.Service.Logger)

	postgresRepo := NewUserPostgresRepository(db, logger)

	columns := []string{"id", "username", "email", "password", "role", "image", "created_at", "updated_at"}

	userUUID := uuid.New()
	userMock := &models.User{
		UserID:   userUUID,
		Username: "test_user",
		Email:    "test@email.com",
		Password: "password",
		Role:     "admin",
		Image:    "avatar.jpeg",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userMock.UserID,
		userMock.Username,
		userMock.Email,
		userMock.Password,
		userMock.Role,
		userMock.Image,
		time.Now(),
		time.Now(),
	)

	mock.ExpectPrepare(findByIDQuery).ExpectQuery().WithArgs(userMock.UserID).WillReturnRows(rows)

	foundUser, err := postgresRepo.FindByID(context.Background(), userMock.UserID)
	require.NoError(t, err)
	require.NotNil(t, foundUser)
	require.Equal(t, foundUser.UserID, userMock.UserID)
}

func TestFindByID_Error(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}
	logger := logger.New(&cfg.Service.Logger)

	postgresRepo := NewUserPostgresRepository(db, logger)

	userUUID := uuid.New()

	mock.ExpectPrepare(findByIDQuery).ExpectQuery().WithArgs(userUUID).WillReturnError(fmt.Errorf("some error"))

	_, err = postgresRepo.FindByID(context.Background(), userUUID)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to find user by ID")
}

func TestFindByID_NotFound(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}
	logger := logger.New(&cfg.Service.Logger)

	postgresRepo := NewUserPostgresRepository(db, logger)

	userUUID := uuid.New()

	mock.ExpectPrepare(findByIDQuery).ExpectQuery().WithArgs(userUUID).WillReturnRows(sqlmock.NewRows(nil))

	_, err = postgresRepo.FindByID(context.Background(), userUUID)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to find user by ID")
}

func TestFindByUsername_Regular(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}
	logger := logger.New(&cfg.Service.Logger)

	postgresRepo := NewUserPostgresRepository(db, logger)

	columns := []string{"id", "username", "email", "password", "role", "image", "created_at", "updated_at"}

	userUUID := uuid.New()
	userMock := &models.User{
		UserID:   userUUID,
		Username: "test_user",
		Email:    "test@email.com",
		Password: "password",
		Role:     "admin",
		Image:    "avatar.jpeg",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userUUID,
		userMock.Username,
		userMock.Email,
		userMock.Password,
		userMock.Role,
		userMock.Image,
		time.Now(),
		time.Now(),
	)

	mock.ExpectPrepare(findByUsernameQuery).
		ExpectQuery().
		WithArgs(userMock.Username).
		WillReturnRows(rows)

	foundUser, err := postgresRepo.FindByUsername(context.Background(), userMock.Username)
	require.NoError(t, err)
	require.NotNil(t, foundUser)
	require.Equal(t, foundUser.UserID, userMock.UserID)
}

func TestFindByUsername_Error(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}
	logger := logger.New(&cfg.Service.Logger)

	postgresRepo := NewUserPostgresRepository(db, logger)

	username := "test_user"

	mock.ExpectPrepare(findByUsernameQuery).
		ExpectQuery().
		WithArgs(username).
		WillReturnError(fmt.Errorf("some error"))

	_, err = postgresRepo.FindByUsername(context.Background(), username)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to find user by username")
}

func TestFindByUsername_NotFound(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}
	logger := logger.New(&cfg.Service.Logger)

	postgresRepo := NewUserPostgresRepository(db, logger)

	username := "test_user"

	mock.ExpectPrepare(findByUsernameQuery).
		ExpectQuery().
		WithArgs(username).
		WillReturnRows(sqlmock.NewRows(nil))

	_, err = postgresRepo.FindByUsername(context.Background(), username)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to find user by username")
}

func TestFindByEmail_Regular(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}
	logger := logger.New(&cfg.Service.Logger)

	postgresRepo := NewUserPostgresRepository(db, logger)

	columns := []string{"id", "username", "email", "password", "role", "image", "created_at", "updated_at"}

	userUUID := uuid.New()
	userMock := &models.User{
		UserID:   userUUID,
		Username: "test_user",
		Email:    "test@email.com",
		Password: "password",
		Role:     "admin",
		Image:    "avatar.jpeg",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userUUID,
		userMock.Username,
		userMock.Email,
		userMock.Password,
		userMock.Role,
		userMock.Image,
		time.Now(),
		time.Now(),
	)

	mock.ExpectPrepare(findByEmailQuery).
		ExpectQuery().
		WithArgs(userMock.Email).
		WillReturnRows(rows)

	foundUser, err := postgresRepo.FindByEmail(context.Background(), userMock.Email)
	require.NoError(t, err)
	require.NotNil(t, foundUser)
	require.Equal(t, foundUser.UserID, userMock.UserID)
}

func TestFindByEmail_Error(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}
	logger := logger.New(&cfg.Service.Logger)

	postgresRepo := NewUserPostgresRepository(db, logger)

	email := "test@email.com"

	mock.ExpectPrepare(findByEmailQuery).
		ExpectQuery().
		WithArgs(email).
		WillReturnError(fmt.Errorf("some error"))

	_, err = postgresRepo.FindByEmail(context.Background(), email)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to find user by email")
}

func TestFindByEmail_NotFound(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}
	logger := logger.New(&cfg.Service.Logger)

	postgresRepo := NewUserPostgresRepository(db, logger)

	email := "test@email.com"

	mock.ExpectPrepare(findByEmailQuery).
		ExpectQuery().
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows(nil))

	_, err = postgresRepo.FindByEmail(context.Background(), email)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to find user by email")
}

func TestUpdate_Regular(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}
	logger := logger.New(&cfg.Service.Logger)

	postgresRepo := NewUserPostgresRepository(db, logger)

	columns := []string{"id", "username", "email", "image", "created_at", "updated_at"}

	userUUID := uuid.New()
	userMock := &models.User{
		UserID:   userUUID,
		Username: "updated_user",
		Email:    "updated@email.com",
		Image:    "new_avatar.jpeg",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userUUID,
		userMock.Username,
		userMock.Email,
		userMock.Image,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(updateUserQuery).WithArgs(
		userMock.Username,
		userMock.Email,
		userMock.Image,
		userMock.UserID,
	).WillReturnRows(rows)

	updatedUser, err := postgresRepo.Update(context.Background(), userMock)
	require.NoError(t, err)
	require.NotNil(t, updatedUser)
	require.Equal(t, userMock.UserID, updatedUser.UserID)
	require.Equal(t, userMock.Username, updatedUser.Username)
	require.Equal(t, userMock.Email, updatedUser.Email)
	require.Equal(t, userMock.Image, updatedUser.Image)
}

func TestUpdate_Error(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}
	logger := logger.New(&cfg.Service.Logger)

	postgresRepo := NewUserPostgresRepository(db, logger)

	userMock := &models.User{
		UserID:   uuid.New(),
		Username: "updated_user",
		Email:    "updated@email.com",
		Image:    "new_avatar.jpeg",
	}

	mock.ExpectQuery(updateUserQuery).WithArgs(
		userMock.Username,
		userMock.Email,
		userMock.Image,
		userMock.UserID,
	).WillReturnError(fmt.Errorf("some error"))

	_, err = postgresRepo.Update(context.Background(), userMock)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to update user")
}

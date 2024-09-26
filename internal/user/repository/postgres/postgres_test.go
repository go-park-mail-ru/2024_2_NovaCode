package postgres

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestInsert_Regular(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	postgresRepo := NewUserPostgresRepository(db)

	columns := []string{"user_id", "username", "email", "password", "role", "created_at", "updated_at"}

	userUUID := uuid.New()
	userMock := &models.User{
		UserID:   userUUID,
		Username: "test_user",
		Email:    "test@email.com",
		Password: "password",
		Role:     "regular",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userUUID,
		userMock.Username,
		userMock.Email,
		userMock.Password,
		userMock.Role,
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

	postgresRepo := NewUserPostgresRepository(db)

	userMock := &models.User{
		Username: "test_user",
		Email:    "test@email.com",
		Password: "password",
		Role:     "regular",
	}

	mock.ExpectQuery(insertUserQuery).WithArgs(
		userMock.Username,
		userMock.Email,
		userMock.Password,
		userMock.Role,
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

	postgresRepo := NewUserPostgresRepository(db)

	columns := []string{"user_id", "username", "email", "password", "role", "created_at", "updated_at"}

	userUUID := uuid.New()
	userMock := &models.User{
		UserID:   userUUID,
		Username: "test_user",
		Email:    "test@email.com",
		Password: "password",
		Role:     "admin",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userUUID,
		userMock.Username,
		userMock.Email,
		userMock.Password,
		userMock.Role,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(findByIDQuery).WithArgs(userMock.UserID).WillReturnRows(rows)

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

	postgresRepo := NewUserPostgresRepository(db)

	userUUID := uuid.New()

	mock.ExpectQuery(findByIDQuery).WithArgs(userUUID).WillReturnError(fmt.Errorf("some error"))

	_, err = postgresRepo.FindByID(context.Background(), userUUID)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to find user by ID")
}

func TestFindByID_NotFound(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	postgresRepo := NewUserPostgresRepository(db)

	userUUID := uuid.New()

	mock.ExpectQuery(findByIDQuery).WithArgs(userUUID).WillReturnRows(sqlmock.NewRows(nil))

	_, err = postgresRepo.FindByID(context.Background(), userUUID)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to find user by ID")
}

func TestFindByUsername_Regular(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	postgresRepo := NewUserPostgresRepository(db)

	columns := []string{"user_id", "username", "email", "password", "role", "created_at", "updated_at"}

	userUUID := uuid.New()
	userMock := &models.User{
		UserID:   userUUID,
		Username: "test_user",
		Email:    "test@email.com",
		Password: "password",
		Role:     "regular",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userUUID,
		userMock.Username,
		userMock.Email,
		userMock.Password,
		userMock.Role,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(findByUsernameQuery).WithArgs(userMock.Username).WillReturnRows(rows)

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

	postgresRepo := NewUserPostgresRepository(db)

	username := "test_user"

	mock.ExpectQuery(findByUsernameQuery).WithArgs(username).WillReturnError(fmt.Errorf("some error"))

	_, err = postgresRepo.FindByUsername(context.Background(), username)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to find user by username")
}

func TestFindByUsername_NotFound(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	postgresRepo := NewUserPostgresRepository(db)

	username := "test_user"

	mock.ExpectQuery(findByUsernameQuery).WithArgs(username).WillReturnRows(sqlmock.NewRows(nil))

	_, err = postgresRepo.FindByUsername(context.Background(), username)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to find user by username")
}

func TestFindByEmail_Regular(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	postgresRepo := NewUserPostgresRepository(db)

	columns := []string{"user_id", "username", "email", "password", "role", "created_at", "updated_at"}

	userUUID := uuid.New()
	userMock := &models.User{
		UserID:   userUUID,
		Username: "test_user",
		Email:    "test@email.com",
		Password: "password",
		Role:     "admin",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userUUID,
		userMock.Username,
		userMock.Email,
		userMock.Password,
		userMock.Role,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(findByEmailQuery).WithArgs(userMock.Email).WillReturnRows(rows)

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

	postgresRepo := NewUserPostgresRepository(db)

	email := "test@email.com"

	mock.ExpectQuery(findByEmailQuery).WithArgs(email).WillReturnError(fmt.Errorf("some error"))

	_, err = postgresRepo.FindByEmail(context.Background(), email)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to find user by email")
}

func TestFindByEmail_NotFound(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	postgresRepo := NewUserPostgresRepository(db)

	email := "test@email.com"

	mock.ExpectQuery(findByEmailQuery).WithArgs(email).WillReturnRows(sqlmock.NewRows(nil))

	_, err = postgresRepo.FindByEmail(context.Background(), email)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to find user by email")
}

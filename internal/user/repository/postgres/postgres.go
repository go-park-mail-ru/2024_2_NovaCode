package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/user"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/google/uuid"
)

const (
	insertUserQuery = `
		INSERT INTO "user" (username, email, password, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id, username, email, password, role, image, created_at, updated_at
	`
	updateUserQuery = `
		UPDATE "user"
		SET username = $1, email = $2, image = $3
		WHERE id = $4
		RETURNING id, username, email, image, created_at, updated_at
	`
	findByIDQuery = `
		SELECT id, username, email, role, password, image, created_at, updated_at
		FROM "user" WHERE id = $1
	`
	findByUsernameQuery = `
		SELECT id, username, email, role, password, image, created_at, updated_at
		FROM "user" WHERE username = $1
	`
	findByEmailQuery = `
		SELECT id, username, email, role, password, image, created_at, updated_at
		FROM "user" WHERE email = $1
	`
)

type UserPostgresRepo struct {
	db     *sql.DB
	logger logger.Logger
}

func NewUserPostgresRepository(db *sql.DB, logger logger.Logger) user.PostgresRepo {
	return &UserPostgresRepo{db, logger}
}

func (repo *UserPostgresRepo) Insert(ctx context.Context, user *models.User) (*models.User, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var insertedUser models.User

	if err := repo.db.QueryRowContext(
		ctx,
		insertUserQuery,
		user.Username,
		user.Email,
		user.Password,
		user.Role,
	).Scan(
		&insertedUser.UserID,
		&insertedUser.Username,
		&insertedUser.Email,
		&insertedUser.Password,
		&insertedUser.Role,
		&insertedUser.Image,
		&insertedUser.CreatedAt,
		&insertedUser.UpdatedAt,
	); err != nil {
		repo.logger.Error(fmt.Sprintf("[user repo] failed Insert: %v", err), requestID)
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}
	repo.logger.Info("[user repo] successful Insert query", requestID)

	return &insertedUser, nil
}

func (repo *UserPostgresRepo) Update(ctx context.Context, user *models.User) (*models.User, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var updatedUser models.User

	if err := repo.db.QueryRowContext(
		ctx,
		updateUserQuery,
		user.Username,
		user.Email,
		user.Image,
		user.UserID,
	).Scan(
		&updatedUser.UserID,
		&updatedUser.Username,
		&updatedUser.Email,
		&updatedUser.Image,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
	); err != nil {
		repo.logger.Error(fmt.Sprintf("[user repo] failed Update: %v", err), requestID)
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	repo.logger.Info("[user repo] successful Update query", requestID)

	return &updatedUser, nil
}

func (repo *UserPostgresRepo) FindByID(ctx context.Context, uuid uuid.UUID) (*models.User, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var user models.User

	if err := repo.db.QueryRowContext(ctx, findByIDQuery, uuid).Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.Image,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		repo.logger.Error(fmt.Sprintf("[user repo] failed FindByID: %v", err), requestID)
		return nil, fmt.Errorf("failed to find user by ID: %w", err)
	}
	repo.logger.Info("[user repo] successful FindByID query", requestID)

	return &user, nil
}

func (repo *UserPostgresRepo) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var user models.User

	if err := repo.db.QueryRowContext(ctx, findByUsernameQuery, username).Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.Image,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		repo.logger.Error(fmt.Sprintf("[user repo] failed FindByUsername: %v", err), requestID)
		return nil, fmt.Errorf("failed to find user by username: %w", err)
	}
	repo.logger.Info("[user repo] successful FindByUsername query", requestID)

	return &user, nil
}

func (repo *UserPostgresRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var user models.User

	if err := repo.db.QueryRowContext(ctx, findByEmailQuery, email).Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.Image,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		repo.logger.Error(fmt.Sprintf("[user repo] failed FindByEmail: %v", err), requestID)
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}
	repo.logger.Info("[user repo] successful FindByEmail query", requestID)

	return &user, nil
}

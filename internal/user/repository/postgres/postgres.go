package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/user"
	"github.com/google/uuid"
)

const (
	insertUserQuery = `
		INSERT INTO "user" (username, email, password, role)
		VALUES ($1, $2, $3, $4)
		RETURNING user_id, username, email, password, role, created_at, updated_at
	`
	findByIDQuery = `
		SELECT user_id, username, email, role, password, created_at, updated_at
		FROM "user" WHERE user_id = $1
	`
	findByUsernameQuery = `
		SELECT user_id, username, email, role, password, created_at, updated_at
		FROM "user" WHERE username = $1
	`
	findByEmailQuery = `
		SELECT user_id, username, email, role, password, created_at, updated_at
		FROM "user" WHERE email = $1
	`
)

type UserPostgresRepo struct {
	db *sql.DB
}

func NewUserPostgresRepository(db *sql.DB) user.Repo {
	return &UserPostgresRepo{db}
}

func (repo *UserPostgresRepo) Insert(ctx context.Context, user *models.User) (*models.User, error) {
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
		&insertedUser.CreatedAt,
		&insertedUser.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	return &insertedUser, nil
}

func (repo *UserPostgresRepo) FindByID(ctx context.Context, uuid uuid.UUID) (*models.User, error) {
	var user models.User

	if err := repo.db.QueryRowContext(ctx, findByIDQuery, uuid).Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to find user by ID: %w", err)
	}

	return &user, nil
}

func (repo *UserPostgresRepo) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	if err := repo.db.QueryRowContext(ctx, findByUsernameQuery, username).Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to find user by username: %w", err)
	}

	return &user, nil
}

func (repo *UserPostgresRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	if err := repo.db.QueryRowContext(ctx, findByEmailQuery, email).Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}

	return &user, nil
}

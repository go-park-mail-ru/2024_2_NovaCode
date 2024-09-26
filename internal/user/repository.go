package user

import (
	"context"

	"github.com/daronenko/auth/internal/models"
	"github.com/google/uuid"
)

type Repo interface {
	Insert(ctx context.Context, user *models.User) (*models.User, error)
	FindByID(ctx context.Context, uuid uuid.UUID) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}

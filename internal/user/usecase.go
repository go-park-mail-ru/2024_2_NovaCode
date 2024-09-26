package user

import (
	"context"

	"github.com/daronenko/auth/internal/models"
	"github.com/daronenko/auth/internal/user/dto"
)

type Usecase interface {
	Register(ctx context.Context, user *models.User) (*dto.UserTokenDTO, error)
	Login(ctx context.Context, user *models.User) (*dto.UserTokenDTO, error)
}

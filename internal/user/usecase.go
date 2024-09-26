package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/user/dto"
)

type Usecase interface {
	Register(ctx context.Context, user *models.User) (*dto.UserTokenDTO, error)
	Login(ctx context.Context, user *models.User) (*dto.UserTokenDTO, error)
}

package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/user/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3"
	"github.com/google/uuid"
)

type Usecase interface {
	Register(ctx context.Context, user *models.User) (*dto.UserTokenDTO, error)
	Login(ctx context.Context, user *models.User) (*dto.UserTokenDTO, error)
	Update(ctx context.Context, user *models.User) (*dto.UserDTO, error)
	UploadImage(ctx context.Context, userID uuid.UUID, file s3.Upload) (*dto.UserDTO, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*dto.UserDTO, error)
}

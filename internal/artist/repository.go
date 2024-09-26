package artist

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

type Repo interface {
	Create(ctx context.Context, artist *models.Artist) (*models.Artist, error)
	FindById(ctx context.Context, artistID uint64) (*models.Artist, error)
	FindByName(ctx context.Context, name string) ([]*models.Artist, error)
}

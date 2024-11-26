package artist

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

type Repo interface {
	Create(ctx context.Context, artist *models.Artist) (*models.Artist, error)
	FindById(ctx context.Context, artistID uint64) (*models.Artist, error)
	GetAll(ctx context.Context) ([]*models.Artist, error)
	FindByQuery(ctx context.Context, query string) ([]*models.Artist, error)
}

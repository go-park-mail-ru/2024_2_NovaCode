package genre

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

type Repo interface {
	Create(ctx context.Context, genre *models.Genre) (*models.Genre, error)
	GetAll(ctx context.Context) ([]*models.Genre, error)
	GetAllByArtistID(ctx context.Context, artistID uint64) ([]*models.Genre, error)
	GetAllByTrackID(ctx context.Context, trackID uint64) ([]*models.Genre, error)
}

package track

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

type Repo interface {
	Create(ctx context.Context, track *models.Track) (*models.Track, error)
	FindById(ctx context.Context, trackID uint64) (*models.Track, error)
	GetAll(ctx context.Context) ([]*models.Track, error)
	GetAllByArtistID(ctx context.Context, artistID int) ([]*models.Track, error)
	FindByName(ctx context.Context, name string) ([]*models.Track, error)
}

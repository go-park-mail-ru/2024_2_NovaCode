package track

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	uuid "github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, track *models.Track) (*models.Track, error)
	FindById(ctx context.Context, trackID uint64) (*models.Track, error)
	GetAll(ctx context.Context) ([]*models.Track, error)
	GetAllByArtistID(ctx context.Context, artistID uint64) ([]*models.Track, error)
	GetAllByAlbumID(ctx context.Context, albumID uint64) ([]*models.Track, error)
	FindByName(ctx context.Context, name string) ([]*models.Track, error)
	AddFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) error
	DeleteFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) error
	IsFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) (bool, error)
	GetFavoriteTracks(ctx context.Context, userID uuid.UUID) ([]*models.Track, error)
}

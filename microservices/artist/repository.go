package artist

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	uuid "github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, artist *models.Artist) (*models.Artist, error)
	FindById(ctx context.Context, artistID uint64) (*models.Artist, error)
	GetAll(ctx context.Context) ([]*models.Artist, error)
	FindByQuery(ctx context.Context, query string) ([]*models.Artist, error)
	AddFavoriteArtist(ctx context.Context, userID uuid.UUID, artistID uint64) error
	DeleteFavoriteArtist(ctx context.Context, userID uuid.UUID, artistID uint64) error
	IsFavoriteArtist(ctx context.Context, userID uuid.UUID, artistID uint64) (bool, error)
	GetFavoriteArtists(ctx context.Context, userID uuid.UUID) ([]*models.Artist, error)
	GetFavoriteArtistsCount(ctx context.Context, userID uuid.UUID) (uint64, error)
	GetArtistLikesCount(ctx context.Context, artistID uint64) (uint64, error)
	GetPopular(ctx context.Context) ([]*models.Artist, error)
}

package album

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	uuid "github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, album *models.Album) (*models.Album, error)
	FindById(ctx context.Context, albumID uint64) (*models.Album, error)
	GetAll(ctx context.Context) ([]*models.Album, error)
	GetAllByArtistID(ctx context.Context, artistID uint64) ([]*models.Album, error)
	FindByQuery(ctx context.Context, query string) ([]*models.Album, error)
	AddFavoriteAlbum(ctx context.Context, userID uuid.UUID, albumID uint64) error
	DeleteFavoriteAlbum(ctx context.Context, userID uuid.UUID, albumID uint64) error
	IsFavoriteAlbum(ctx context.Context, userID uuid.UUID, albumID uint64) (bool, error)
	GetFavoriteAlbums(ctx context.Context, userID uuid.UUID) ([]*models.Album, error)
	GetFavoriteAlbumsCount(ctx context.Context, userID uuid.UUID) (uint64, error)
	GetAlbumLikesCount(ctx context.Context, albumID uint64) (uint64, error)
}

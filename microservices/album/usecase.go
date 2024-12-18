package album

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/dto"
	uuid "github.com/google/uuid"
)

type Usecase interface {
	View(ctx context.Context, albumID uint64) (*dto.AlbumDTO, error)
	Search(ctx context.Context, name string) ([]*dto.AlbumDTO, error)
	GetAll(ctx context.Context) ([]*dto.AlbumDTO, error)
	GetAllByArtistID(ctx context.Context, artistID uint64) ([]*dto.AlbumDTO, error)
	AddFavoriteAlbum(ctx context.Context, userID uuid.UUID, albumID uint64) error
	DeleteFavoriteAlbum(ctx context.Context, userID uuid.UUID, albumID uint64) error
	IsFavoriteAlbum(ctx context.Context, userID uuid.UUID, albumID uint64) (bool, error)
	GetFavoriteAlbums(ctx context.Context, userID uuid.UUID) ([]*dto.AlbumDTO, error)
	GetFavoriteAlbumsCount(ctx context.Context, userID uuid.UUID) (uint64, error)
	GetAlbumLikesCount(ctx context.Context, albumID uint64) (uint64, error)
}

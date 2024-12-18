package artist

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/dto"
	uuid "github.com/google/uuid"
)

type Usecase interface {
	View(ctx context.Context, artistID uint64) (*dto.ArtistDTO, error)
	Search(ctx context.Context, query string) ([]*dto.ArtistDTO, error)
	GetAll(ctx context.Context) ([]*dto.ArtistDTO, error)
	AddFavoriteArtist(ctx context.Context, userID uuid.UUID, artistID uint64) error
	DeleteFavoriteArtist(ctx context.Context, userID uuid.UUID, artistID uint64) error
	IsFavoriteArtist(ctx context.Context, userID uuid.UUID, artistID uint64) (bool, error)
	GetFavoriteArtists(ctx context.Context, userID uuid.UUID) ([]*dto.ArtistDTO, error)
	GetPopular(ctx context.Context) ([]*dto.ArtistDTO, error)
}

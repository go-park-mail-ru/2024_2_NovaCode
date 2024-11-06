package genre

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/genre/dto"
)

type Usecase interface {
	GetAll(ctx context.Context) ([]*dto.GenreDTO, error)
	GetAllByArtistID(ctx context.Context, artistID uint64) ([]*dto.GenreDTO, error)
	GetAllByAlbumID(ctx context.Context, albumID uint64) ([]*dto.GenreDTO, error)
	GetAllByTrackID(ctx context.Context, trackID uint64) ([]*dto.GenreDTO, error)
}

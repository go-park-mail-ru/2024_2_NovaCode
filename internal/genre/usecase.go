package genre

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/genre/dto"
)

type Usecase interface {
	GetAll(ctx context.Context) ([]*dto.GenreDTO, error)
	GetAllByArtistID(ctx context.Context, artistID int) ([]*dto.GenreDTO, error)
	GetAllByAlbumID(ctx context.Context, albumID int) ([]*dto.GenreDTO, error)
	GetAllByTrackID(ctx context.Context, trackID int) ([]*dto.GenreDTO, error)
}

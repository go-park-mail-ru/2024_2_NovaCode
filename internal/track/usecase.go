package track

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/track/dto"
)

type Usecase interface {
	View(ctx context.Context, trackID uint64) (*dto.TrackDTO, error)
	Search(ctx context.Context, name string) ([]*dto.TrackDTO, error)
	GetAll(ctx context.Context) ([]*dto.TrackDTO, error)
	GetAllByArtistID(ctx context.Context, artistID uint64) ([]*dto.TrackDTO, error)
}

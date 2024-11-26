package artist

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/dto"
)

type Usecase interface {
	View(ctx context.Context, artistID uint64) (*dto.ArtistDTO, error)
	Search(ctx context.Context, query string) ([]*dto.ArtistDTO, error)
	GetAll(ctx context.Context) ([]*dto.ArtistDTO, error)
}

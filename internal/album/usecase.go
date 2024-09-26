package album

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/album/dto"
)

type Usecase interface {
	View(ctx context.Context, albumID uint64) (*dto.AlbumDTO, error)
	Search(ctx context.Context, name string) ([]*dto.AlbumDTO, error)
}

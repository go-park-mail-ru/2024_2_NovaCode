package album

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

type Repo interface {
	Create(ctx context.Context, album *models.Album) (*models.Album, error)
	FindById(ctx context.Context, albumID uint64) (*models.Album, error)
	GetAll(ctx context.Context) ([]*models.Album, error)
	GetAllByArtistID(ctx context.Context, artistID uint64) ([]*models.Album, error)
	FindByName(ctx context.Context, name string) ([]*models.Album, error)
}

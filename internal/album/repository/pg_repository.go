package repository

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

type AlbumRepository struct {
	db *sql.DB
}

func NewAlbumPGRepository(db *sql.DB) *AlbumRepository {
	return &AlbumRepository{db: db}
}

func (r *AlbumRepository) Create(ctx context.Context, album *models.Album) (*models.Album, error) {
	return nil, nil
}

func (r *AlbumRepository) FindById(ctx context.Context, albumID uint64) (*models.Album, error) {
	return nil, nil
}

func (r *AlbumRepository) FindByName(ctx context.Context, name string) (*models.Album, error) {
	return nil, nil
}

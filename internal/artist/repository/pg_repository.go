package repository

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

type ArtistRepository struct {
	db *sql.DB
}

func NewArtistPGRepository(db *sql.DB) *ArtistRepository {
	return &ArtistRepository{db: db}
}

func (r *ArtistRepository) Create(ctx context.Context, artist *models.Artist) (*models.Artist, error) {
	return nil, nil
}

func (r *ArtistRepository) FindById(ctx context.Context, artistID uint64) (*models.Artist, error) {
	return nil, nil
}

func (r *ArtistRepository) FindByName(ctx context.Context, name string) (*models.Artist, error) {
	return nil, nil
}

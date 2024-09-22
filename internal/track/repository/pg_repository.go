package repository

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

type TrackRepository struct {
	db *sql.DB
}

func NewTrackPGRepository(db *sql.DB) *TrackRepository {
	return &TrackRepository{db: db}
}

func (r *TrackRepository) Create(ctx context.Context, track *models.Track) (*models.Track, error) {
	return nil, nil
}

func (r *TrackRepository) FindById(ctx context.Context, trackID uint64) (*models.Track, error) {
	return nil, nil
}

func (r *TrackRepository) FindByName(ctx context.Context, name string) (*models.Track, error) {
	return nil, nil
}

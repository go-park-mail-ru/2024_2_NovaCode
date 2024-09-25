package repository

import (
	"context"
	"database/sql"
	"strings"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/pkg/errors"
)

type TrackRepository struct {
	db *sql.DB
}

func NewTrackPGRepository(db *sql.DB) *TrackRepository {
	return &TrackRepository{db: db}
}

type TrackRepositoryInterface interface {
	Create(ctx context.Context, track *models.Track) (*models.Track, error)
	UpdateTrack(ctx context.Context, track *models.Track) error
	FindById(ctx context.Context, trackID uint64) (*models.Track, error)
	FindByName(ctx context.Context, name string) ([]*models.Track, error)
}

func (r *TrackRepository) Create(ctx context.Context, track *models.Track) (*models.Track, error) {
	createdTrack := &models.Track{}
	row := r.db.QueryRowContext(
		ctx,
		createTrackQuery,
		track.Name,
		track.Genre,
		track.Duration,
		track.FilePath,
		track.Image,
		track.ArtistID,
		track.AlbumID,
		track.ReleaseDate,
	)

	if err := row.Scan(
		&createdTrack.ID,
		&createdTrack.Name,
		&createdTrack.Genre,
		&createdTrack.Duration,
		&createdTrack.FilePath,
		&createdTrack.Image,
		&createdTrack.ArtistID,
		&createdTrack.AlbumID,
		&createdTrack.ReleaseDate,
		&createdTrack.CreatedAt,
		&createdTrack.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "Create.Query")
	}

	return createdTrack, nil
}

func (r *TrackRepository) FindById(ctx context.Context, trackID uint64) (*models.Track, error) {
	track := &models.Track{}
	row := r.db.QueryRowContext(ctx, findByIDQuery, trackID)
	if err := row.Scan(
		&track.ID,
		&track.Name,
		&track.Genre,
		&track.Duration,
		&track.FilePath,
		&track.Image,
		&track.ArtistID,
		&track.AlbumID,
		&track.ReleaseDate,
		&track.CreatedAt,
		&track.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "FindById.Query")
	}

	return track, nil
}

func (r *TrackRepository) FindByName(ctx context.Context, name string) ([]*models.Track, error) {
	var tracks []*models.Track
	rows, err := r.db.QueryContext(ctx, findByNameQuery, name)
	if err != nil {
		return nil, errors.Wrap(err, "FindByName.Query")
	}
	defer rows.Close()

	for rows.Next() {
		track := &models.Track{}
		err := rows.Scan(
			&track.ID,
			&track.Name,
			&track.Genre,
			&track.Duration,
			&track.FilePath,
			&track.Image,
			&track.ArtistID,
			&track.AlbumID,
			&track.ReleaseDate,
			&track.CreatedAt,
			&track.UpdatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "FindByName.Rows.Scan")
		}
		if strings.Contains(track.Name, name) {
			tracks = append(tracks, track)
		}
	}

	return tracks, nil
}

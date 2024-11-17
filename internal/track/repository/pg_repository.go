package repository

import (
	"context"
	"database/sql"
	"strings"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	uuid "github.com/google/uuid"
	"github.com/pkg/errors"
)

type TrackRepository struct {
	db *sql.DB
}

func NewTrackPGRepository(db *sql.DB) *TrackRepository {
	return &TrackRepository{db: db}
}

func (r *TrackRepository) Create(ctx context.Context, track *models.Track) (*models.Track, error) {
	createdTrack := &models.Track{}
	row := r.db.QueryRowContext(
		ctx,
		createTrackQuery,
		track.Name,
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
			return nil, errors.Wrap(err, "FindByName.Query")
		}
		if strings.Contains(track.Name, name) {
			tracks = append(tracks, track)
		}
	}

	return tracks, nil
}

func (r *TrackRepository) GetAll(ctx context.Context) ([]*models.Track, error) {
	var tracks []*models.Track
	rows, err := r.db.QueryContext(ctx, getAllQuery)
	if err != nil {
		return nil, errors.Wrap(err, "GetAll.Query")
	}
	defer rows.Close()

	for rows.Next() {
		track := &models.Track{}
		err := rows.Scan(
			&track.ID,
			&track.Name,
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
			return nil, errors.Wrap(err, "GetAll.Query")
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (r *TrackRepository) GetAllByArtistID(ctx context.Context, artistID uint64) ([]*models.Track, error) {
	var tracks []*models.Track
	rows, err := r.db.QueryContext(ctx, getByArtistIDQuery, artistID)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllByArtistID.Query")
	}
	defer rows.Close()

	for rows.Next() {
		track := &models.Track{}
		err := rows.Scan(
			&track.ID,
			&track.Name,
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
			return nil, errors.Wrap(err, "GetAllByArtistID.Query")
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (r *TrackRepository) AddFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) error {
	_, err := r.db.ExecContext(ctx, addFavoriteTrackQuery, userID, trackID)
	if err != nil {
		return errors.Wrap(err, "AddFavoriteTrack.Query")
	}

	return nil
}

func (r *TrackRepository) DeleteFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) error {
	_, err := r.db.ExecContext(ctx, deleteFavoriteTrackQuery, userID, trackID)
	if err != nil {
		return errors.Wrap(err, "DeleteFavoriteTrack.Query")
	}

	return nil
}

func (r *TrackRepository) IsFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx, isFavoriteTrackQuery, userID, trackID).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.Wrap(err, "IsFavoriteTrack.Query")
	}

	return exists, nil
}

func (r *TrackRepository) GetFavoriteTracks(ctx context.Context, userID uuid.UUID) ([]*models.Track, error) {
	var tracks []*models.Track
	rows, err := r.db.QueryContext(ctx, getFavoriteQuery, userID)
	if err != nil {
		return nil, errors.Wrap(err, "GetFavoriteTracks.Query")
	}
	defer rows.Close()

	for rows.Next() {
		track := &models.Track{}
		err := rows.Scan(
			&track.ID,
			&track.Name,
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
			return nil, errors.Wrap(err, "GetFavoriteTracks.Query")
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

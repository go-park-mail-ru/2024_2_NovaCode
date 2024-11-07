package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/pkg/errors"
)

type TrackRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewTrackPGRepository(db *sql.DB, logger logger.Logger) *TrackRepository {
	return &TrackRepository{db, logger}
}

func (r *TrackRepository) Create(ctx context.Context, track *models.Track) (*models.Track, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
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
		r.logger.Error(fmt.Sprintf("[track repo] failed to scan columns in Create: %v", err), requestID)
		return nil, errors.Wrap(err, "Create.Query")
	}
	r.logger.Info("[track repo] successful Create query", requestID)

	return createdTrack, nil
}

func (r *TrackRepository) FindById(ctx context.Context, trackID uint64) (*models.Track, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
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
		r.logger.Error(fmt.Sprintf("[track repo] failed to scan columns in FindById: %v", err), requestID)
		return nil, errors.Wrap(err, "FindById.Query")
	}
	r.logger.Info("[track repo] successful FindById query", requestID)

	return track, nil
}

func (r *TrackRepository) FindByName(ctx context.Context, name string) ([]*models.Track, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var tracks []*models.Track
	rows, err := r.db.QueryContext(ctx, findByNameQuery, name)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to execute FindByName query: %v", err), requestID)
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
			r.logger.Error(fmt.Sprintf("[track repo] failed to scan columns in FindByName: %v", err), requestID)
			return nil, errors.Wrap(err, "FindByName.Query")
		}
		if strings.Contains(track.Name, name) {
			tracks = append(tracks, track)
		}
	}
	r.logger.Info("[track repo] successful FindByName query", requestID)

	return tracks, nil
}

func (r *TrackRepository) GetAll(ctx context.Context) ([]*models.Track, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var tracks []*models.Track
	rows, err := r.db.QueryContext(ctx, getAllQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to execute GetAll query: %v", err), requestID)
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
			r.logger.Error(fmt.Sprintf("[track repo] failed to scan columns in GetAll: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAll.Query")
		}
		tracks = append(tracks, track)
	}
	r.logger.Info("[track repo] successful GetAll query", requestID)

	return tracks, nil
}

func (r *TrackRepository) GetAllByArtistID(ctx context.Context, artistID uint64) ([]*models.Track, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var tracks []*models.Track
	rows, err := r.db.QueryContext(ctx, getByArtistIDQuery, artistID)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to execute GetAllByArtistID query: %v", err), requestID)
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
			r.logger.Error(fmt.Sprintf("[track repo] failed to scan columns in GetAllByArtistID: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAllByArtistID.Query")
		}
		tracks = append(tracks, track)
	}
	r.logger.Info("[track repo] successful GetAllByArtistID query", requestID)

	return tracks, nil
}

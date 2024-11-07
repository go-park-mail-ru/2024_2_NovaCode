package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/pkg/errors"
)

type GenreRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewGenrePGRepository(db *sql.DB, logger logger.Logger) *GenreRepository {
	return &GenreRepository{db, logger}
}

func (r *GenreRepository) Create(ctx context.Context, genre *models.Genre) (*models.Genre, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	createdGenre := &models.Genre{}

	row := r.db.QueryRowContext(
		ctx,
		createGenreQuery,
		genre.Name,
		genre.RusName,
	)

	if err := row.Scan(
		&createdGenre.ID,
		&createdGenre.Name,
		&createdGenre.RusName,
		&createdGenre.CreatedAt,
		&createdGenre.UpdatedAt,
	); err != nil {
		r.logger.Error(fmt.Sprintf("[genre repo] failed to scan columns in Create: %v", err), requestID)
		return nil, errors.Wrap(err, "Create.Query")
	}
	r.logger.Info("[genre repo] successful Create query", requestID)

	return createdGenre, nil
}

func (r *GenreRepository) FindById(ctx context.Context, genreID uint64) (*models.Genre, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	genre := &models.Genre{}
	row := r.db.QueryRowContext(ctx, findByIDQuery, genreID)
	if err := row.Scan(
		&genre.ID,
		&genre.Name,
		&genre.RusName,
		&genre.CreatedAt,
		&genre.UpdatedAt,
	); err != nil {
		r.logger.Error(fmt.Sprintf("[genre repo] failed to scan columns in FindById: %v", err), requestID)
		return nil, errors.Wrap(err, "FindById.Query")
	}
	r.logger.Info("[genre repo] successful FindById query", requestID)

	return genre, nil
}

func (r *GenreRepository) GetAll(ctx context.Context) ([]*models.Genre, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var genres []*models.Genre
	rows, err := r.db.QueryContext(ctx, getAllQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[genre repo] failed to execute GetAll query: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAll.Query")
	}
	defer rows.Close()

	for rows.Next() {
		genre := &models.Genre{}
		err := rows.Scan(
			&genre.ID,
			&genre.Name,
			&genre.RusName,
			&genre.CreatedAt,
			&genre.UpdatedAt,
		)
		if err != nil {
			r.logger.Error(fmt.Sprintf("[genre repo] failed to scan columns in GetAll: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAll.Scan")
		}
		genres = append(genres, genre)
	}
	r.logger.Info("[genre repo] successful GetAll query", requestID)

	return genres, nil
}

func (r *GenreRepository) GetAllByArtistID(ctx context.Context, artistID uint64) ([]*models.Genre, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var genres []*models.Genre
	rows, err := r.db.QueryContext(ctx, getByArtistIDQuery, artistID)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[genre repo] failed to execute GetAllByArtistID query: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAllByArtistID.Query")
	}
	defer rows.Close()

	for rows.Next() {
		genre := &models.Genre{}
		err := rows.Scan(
			&genre.ID,
			&genre.Name,
			&genre.RusName,
			&genre.CreatedAt,
			&genre.UpdatedAt,
		)
		if err != nil {
			r.logger.Error(fmt.Sprintf("[genre repo] failed to scan columns in GetAllByArtistID: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAllByArtistID.Scan")
		}
		genres = append(genres, genre)
	}
	r.logger.Info("[genre repo] successful GetAllByArtistID query", requestID)

	return genres, nil
}

func (r *GenreRepository) GetAllByAlbumID(ctx context.Context, albumID uint64) ([]*models.Genre, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var genres []*models.Genre
	rows, err := r.db.QueryContext(ctx, getByAlbumIDQuery, albumID)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[genre repo] failed to execute GetAllByAlbumID query: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAllByAlbumID.Query")
	}
	defer rows.Close()

	for rows.Next() {
		genre := &models.Genre{}
		err := rows.Scan(
			&genre.ID,
			&genre.Name,
			&genre.RusName,
			&genre.CreatedAt,
			&genre.UpdatedAt,
		)
		if err != nil {
			r.logger.Error(fmt.Sprintf("[genre repo] failed to scan columns in GetAllByAlbumID: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAllByAlbumID.Scan")
		}
		genres = append(genres, genre)
	}
	r.logger.Info("[genre repo] successful GetAllByAlbumID query", requestID)

	return genres, nil
}

func (r *GenreRepository) GetAllByTrackID(ctx context.Context, trackID uint64) ([]*models.Genre, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var genres []*models.Genre
	rows, err := r.db.QueryContext(ctx, getByTrackIDQuery, trackID)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[genre repo] failed to execute GetAllByTrackID query: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAllByTrackID.Query")
	}
	defer rows.Close()

	for rows.Next() {
		genre := &models.Genre{}
		err := rows.Scan(
			&genre.ID,
			&genre.Name,
			&genre.RusName,
			&genre.CreatedAt,
			&genre.UpdatedAt,
		)
		if err != nil {
			r.logger.Error(fmt.Sprintf("[genre repo] failed to scan columns in GetAllByTrackID: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAllByTrackID.Scan")
		}
		genres = append(genres, genre)
	}
	r.logger.Info("[genre repo] successful GetAllByTrackID query", requestID)

	return genres, nil
}

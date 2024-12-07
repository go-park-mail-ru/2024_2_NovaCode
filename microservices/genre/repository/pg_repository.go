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
		r.logger.Error(fmt.Sprintf("[genre repo] failed to scan row in Create: %v", err), requestID)
		return nil, errors.Wrap(err, "Create.Scan")
	}
	r.logger.Info("[genre repo] successful scan row", requestID)

	return createdGenre, nil
}

func (r *GenreRepository) FindById(ctx context.Context, genreID uint64) (*models.Genre, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	genre := &models.Genre{}

	stmt, err := r.db.PrepareContext(ctx, findByIDQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[genre repo] failed to prepare context in FindById: %v", err), requestID)
		return nil, errors.Wrap(err, "FindById.PrepareContext")
	}
	r.logger.Info("[genre repo] successful FindById prepare context", requestID)
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, genreID)
	if err := row.Scan(
		&genre.ID,
		&genre.Name,
		&genre.RusName,
		&genre.CreatedAt,
		&genre.UpdatedAt,
	); err != nil {
		r.logger.Error(fmt.Sprintf("[genre repo] failed to scan row in FindById: %v", err), requestID)
		return nil, errors.Wrap(err, "FindById.Scan")
	}
	r.logger.Info("[genre repo] successful FindById scan row", requestID)

	return genre, nil
}

func (r *GenreRepository) GetAll(ctx context.Context) ([]*models.Genre, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var genres []*models.Genre
	rows, err := r.db.QueryContext(ctx, getAllQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[genre repo] failed to query context in GetAll: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAll.Query")
	}
	r.logger.Info("[genre repo] successful GetAll query context", requestID)
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
			r.logger.Error(fmt.Sprintf("[genre repo] failed to scan rows in GetAll: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAll.Scan")
		}
		r.logger.Info("[genre repo] successful GetAll scan rows", requestID)

		genres = append(genres, genre)
	}

	return genres, nil
}

func (r *GenreRepository) GetAllByArtistID(ctx context.Context, artistID uint64) ([]*models.Genre, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	stmt, err := r.db.PrepareContext(ctx, getByArtistIDQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[genre repo] failed to prepare context in GetAllByArtistID: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAllByArtistID.PrepareContext")
	}
	r.logger.Info("[genre repo] successful GetAllByArtistID prepare context", requestID)
	defer stmt.Close()

	var genres []*models.Genre
	rows, err := stmt.QueryContext(ctx, artistID)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[genre repo] failed to query context in GetAllByArtistID: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAllByArtistID.QueryContext")
	}
	r.logger.Info("[genre repo] successful GetAllByArtistID query context", requestID)
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
			r.logger.Error(fmt.Sprintf("[genre repo] failed to scan rows in GetAllByArtistID: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAllByArtistID.Scan")
		}
		r.logger.Info("[genre repo] successful GetAllByArtistID scan rows", requestID)

		genres = append(genres, genre)
	}

	return genres, nil
}

func (r *GenreRepository) GetAllByTrackID(ctx context.Context, trackID uint64) ([]*models.Genre, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	stmt, err := r.db.PrepareContext(ctx, getByTrackIDQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[genre repo] failed to prepare context in GetAllByTrackID: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAllByTrackID.PrepareContext")
	}
	r.logger.Info("[genre repo] successful GetAllByTrackID prepare context", requestID)
	defer stmt.Close()

	var genres []*models.Genre
	rows, err := stmt.QueryContext(ctx, trackID)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[genre repo] failed to query context in GetAllByTrackID: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAllByTrackID.QueryContext")
	}
	r.logger.Info("[genre repo] successful GetAllByTrackID query context", requestID)
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
			r.logger.Error(fmt.Sprintf("[genre repo] failed to scan rows in GetAllByTrackID: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAllByTrackID.Scan")
		}
		r.logger.Info("[genre repo] successful GetAllByTrackID scan rows", requestID)

		genres = append(genres, genre)
	}

	return genres, nil
}

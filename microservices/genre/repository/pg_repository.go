package repository

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/pkg/errors"
)

type GenreRepository struct {
	db *sql.DB
}

func NewGenrePGRepository(db *sql.DB) *GenreRepository {
	return &GenreRepository{db: db}
}

func (r *GenreRepository) Create(ctx context.Context, genre *models.Genre) (*models.Genre, error) {
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
		return nil, errors.Wrap(err, "Create.Query")
	}

	return createdGenre, nil
}

func (r *GenreRepository) FindById(ctx context.Context, genreID uint64) (*models.Genre, error) {
	genre := &models.Genre{}

	stmt, err := r.db.PrepareContext(ctx, findByIDQuery)
	if err != nil {
		return nil, errors.Wrap(err, "FindById.PrepareContext")
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, genreID)
	if err := row.Scan(
		&genre.ID,
		&genre.Name,
		&genre.RusName,
		&genre.CreatedAt,
		&genre.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "FindById.QueryRow")
	}

	return genre, nil
}

func (r *GenreRepository) GetAll(ctx context.Context) ([]*models.Genre, error) {
	var genres []*models.Genre
	rows, err := r.db.QueryContext(ctx, getAllQuery)
	if err != nil {
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
			return nil, errors.Wrap(err, "GetAll.Scan")
		}
		genres = append(genres, genre)
	}

	return genres, nil
}

func (r *GenreRepository) GetAllByArtistID(ctx context.Context, artistID uint64) ([]*models.Genre, error) {
	stmt, err := r.db.PrepareContext(ctx, getByArtistIDQuery)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllByArtistID.PrepareContext")
	}
	defer stmt.Close()

	var genres []*models.Genre
	rows, err := stmt.QueryContext(ctx, artistID)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllByArtistID.QueryContext")
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
			return nil, errors.Wrap(err, "GetAllByArtistID.Scan")
		}
		genres = append(genres, genre)
	}

	return genres, nil
}

func (r *GenreRepository) GetAllByTrackID(ctx context.Context, trackID uint64) ([]*models.Genre, error) {
	stmt, err := r.db.PrepareContext(ctx, getByTrackIDQuery)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllByTrackID.PrepareContext")
	}
	defer stmt.Close()

	var genres []*models.Genre
	rows, err := stmt.QueryContext(ctx, trackID)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllByTrackID.QueryContext")
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
			return nil, errors.Wrap(err, "GetAllByTrackID.Scan")
		}
		genres = append(genres, genre)
	}

	return genres, nil
}

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

// GetAll retrieves all genres from the database.
func (r *GenreRepository) GetAll(ctx context.Context) ([]*models.Genre, error) {
	var genres []*models.Genre
	rows, err := r.db.QueryContext(ctx, getAllGenresQuery)
	if err != nil {
		return nil, errors.Wrap(err, "GetAll.Query")
	}
	defer rows.Close()

	for rows.Next() {
		genre := &models.Genre{}
		err := rows.Scan(&genre.ID, &genre.Name, &genre.RusName)
		if err != nil {
			return nil, errors.Wrap(err, "GetAll.Scan")
		}
		genres = append(genres, genre)
	}

	return genres, nil
}

// GetAllByArtistID retrieves all genres associated with a specific artist ID.
func (r *GenreRepository) GetAllByArtistID(ctx context.Context, artistID int) ([]*models.Genre, error) {
	var genres []*models.Genre
	rows, err := r.db.QueryContext(ctx, getAllGenresByArtistIDQuery, artistID)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllByArtistID.Query")
	}
	defer rows.Close()

	for rows.Next() {
		genre := &models.Genre{}
		err := rows.Scan(&genre.ID, &genre.Name, &genre.RusName)
		if err != nil {
			return nil, errors.Wrap(err, "GetAllByArtistID.Scan")
		}
		genres = append(genres, genre)
	}

	return genres, nil
}

// GetAllByAlbumID retrieves all genres associated with a specific album ID.
func (r *GenreRepository) GetAllByAlbumID(ctx context.Context, albumID int) ([]*models.Genre, error) {
	var genres []*models.Genre
	rows, err := r.db.QueryContext(ctx, getAllGenresByAlbumIDQuery, albumID)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllByAlbumID.Query")
	}
	defer rows.Close()

	for rows.Next() {
		genre := &models.Genre{}
		err := rows.Scan(&genre.ID, &genre.Name, &genre.RusName)
		if err != nil {
			return nil, errors.Wrap(err, "GetAllByAlbumID.Scan")
		}
		genres = append(genres, genre)
	}

	return genres, nil
}

// GetAllByTrackID retrieves all genres associated with a specific track ID.
func (r *GenreRepository) GetAllByTrackID(ctx context.Context, trackID int) ([]*models.Genre, error) {
	var genres []*models.Genre
	rows, err := r.db.QueryContext(ctx, getAllGenresByTrackIDQuery, trackID)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllByTrackID.Query")
	}
	defer rows.Close()

	for rows.Next() {
		genre := &models.Genre{}
		err := rows.Scan(&genre.ID, &genre.Name, &genre.RusName)
		if err != nil {
			return nil, errors.Wrap(err, "GetAllByTrackID.Scan")
		}
		genres = append(genres, genre)
	}

	return genres, nil
}

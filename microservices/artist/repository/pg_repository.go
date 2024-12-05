package repository

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/pkg/errors"
)

type ArtistRepository struct {
	db *sql.DB
}

func NewArtistPGRepository(db *sql.DB) *ArtistRepository {
	return &ArtistRepository{db: db}
}

func (r *ArtistRepository) Create(ctx context.Context, artist *models.Artist) (*models.Artist, error) {
	createdArtist := &models.Artist{}
	row := r.db.QueryRowContext(
		ctx,
		createArtistQuery,
		artist.Name,
		artist.Bio,
		artist.Country,
		artist.Image,
	)

	if err := row.Scan(
		&createdArtist.ID,
		&createdArtist.Name,
		&createdArtist.Bio,
		&createdArtist.Country,
		&createdArtist.Image,
		&createdArtist.CreatedAt,
		&createdArtist.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "Create.Query")
	}

	return createdArtist, nil
}

func (r *ArtistRepository) FindById(ctx context.Context, artistID uint64) (*models.Artist, error) {
	artist := &models.Artist{}

	stmt, err := r.db.PrepareContext(ctx, findByIDQuery)
	if err != nil {
		return nil, errors.Wrap(err, "FindById.PrepareContext")
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, artistID)
	if err := row.Scan(
		&artist.ID,
		&artist.Name,
		&artist.Bio,
		&artist.Country,
		&artist.Image,
		&artist.CreatedAt,
		&artist.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "FindById.QueryRow")
	}

	return artist, nil
}

func (r *ArtistRepository) FindByQuery(ctx context.Context, query string) ([]*models.Artist, error) {
	tsQuery := utils.MakeSearchQuery(query)

	stmt, err := r.db.PrepareContext(ctx, findByQuery)
	if err != nil {
		return nil, errors.Wrap(err, "FindByQuery.PrepareContext")
	}
	defer stmt.Close()

	var artists []*models.Artist
	rows, err := stmt.QueryContext(ctx, tsQuery)
	if err != nil {
		return nil, errors.Wrap(err, "FindByQuery.QueryContext")
	}
	defer rows.Close()

	for rows.Next() {
		artist := &models.Artist{}
		err := rows.Scan(
			&artist.ID,
			&artist.Name,
			&artist.Bio,
			&artist.Country,
			&artist.Image,
			&artist.CreatedAt,
			&artist.UpdatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "FindByQuery.Scan")
		}

		artists = append(artists, artist)
	}

	return artists, nil
}

func (r *ArtistRepository) GetAll(ctx context.Context) ([]*models.Artist, error) {
	var artists []*models.Artist
	rows, err := r.db.QueryContext(ctx, getAllQuery)
	if err != nil {
		return nil, errors.Wrap(err, "GetAll.Query")
	}
	defer rows.Close()

	for rows.Next() {
		artist := &models.Artist{}
		err := rows.Scan(
			&artist.ID,
			&artist.Name,
			&artist.Bio,
			&artist.Country,
			&artist.Image,
			&artist.CreatedAt,
			&artist.UpdatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "GetAll.Query")
		}
		artists = append(artists, artist)
	}

	return artists, nil
}

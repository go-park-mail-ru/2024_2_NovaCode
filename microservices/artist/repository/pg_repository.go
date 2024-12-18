package repository

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	uuid "github.com/google/uuid"
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
	row := r.db.QueryRowContext(ctx, findByIDQuery, artistID)
	if err := row.Scan(
		&artist.ID,
		&artist.Name,
		&artist.Bio,
		&artist.Country,
		&artist.Image,
		&artist.CreatedAt,
		&artist.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "FindById.Query")
	}

	return artist, nil
}

func (r *ArtistRepository) FindByQuery(ctx context.Context, query string) ([]*models.Artist, error) {
	tsQuery := utils.MakeSearchQuery(query)

	var artists []*models.Artist
	rows, err := r.db.QueryContext(ctx, findByQuery, tsQuery)
	if err != nil {
		return nil, errors.Wrap(err, "FindByQuery.Query")
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
			return nil, errors.Wrap(err, "FindByQuery.Query")
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

func (r *ArtistRepository) AddFavoriteArtist(ctx context.Context, userID uuid.UUID, artistID uint64) error {
	_, err := r.db.ExecContext(ctx, addFavoriteArtistQuery, userID, artistID)
	if err != nil {
		return errors.Wrap(err, "AddFavoriteArtist.Query")
	}

	return nil
}

func (r *ArtistRepository) DeleteFavoriteArtist(ctx context.Context, userID uuid.UUID, artistID uint64) error {
	_, err := r.db.ExecContext(ctx, deleteFavoriteArtistQuery, userID, artistID)
	if err != nil {
		return errors.Wrap(err, "DeleteFavoriteArtist.Query")
	}

	return nil
}

func (r *ArtistRepository) IsFavoriteArtist(ctx context.Context, userID uuid.UUID, artistID uint64) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx, isFavoriteArtistQuery, userID, artistID).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.Wrap(err, "IsFavoriteArtist.Query")
	}

	return exists, nil
}

func (r *ArtistRepository) GetFavoriteArtists(ctx context.Context, userID uuid.UUID) ([]*models.Artist, error) {
	var artists []*models.Artist
	rows, err := r.db.QueryContext(ctx, getFavoriteQuery, userID)
	if err != nil {
		return nil, errors.Wrap(err, "GetFavoriteArtists.Query")
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
			return nil, errors.Wrap(err, "GetFavoriteArtists.Query")
		}
		artists = append(artists, artist)
	}

	return artists, nil
}

func (r *ArtistRepository) GetFavoriteArtistsCount(ctx context.Context, userID uuid.UUID) (uint64, error) {
	var count uint64
	err := r.db.QueryRowContext(ctx, getFavoriteCountQuery, userID).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return 0, errors.Wrap(err, "GetFavoriteArtistsCount.Query")
	}

	return count, nil
}

func (r *ArtistRepository) GetArtistLikesCount(ctx context.Context, artistID uint64) (uint64, error) {
	var likesCount uint64
	err := r.db.QueryRowContext(ctx, getLikesCountQuery, artistID).Scan(&likesCount)
	if err != nil && err != sql.ErrNoRows {
		return 0, errors.Wrap(err, "GetLikesCount.Query")
	}

	return likesCount, nil
}

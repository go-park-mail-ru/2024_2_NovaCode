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

type ArtistRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewArtistPGRepository(db *sql.DB, logger logger.Logger) *ArtistRepository {
	return &ArtistRepository{db, logger}
}

func (r *ArtistRepository) Create(ctx context.Context, artist *models.Artist) (*models.Artist, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
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
		r.logger.Error(fmt.Sprintf("[artist repo] failed to scan columns in Create: %v", err), requestID)
		return nil, errors.Wrap(err, "Create.Query")
	}
	r.logger.Info("[artist repo] successful Create query", requestID)

	return createdArtist, nil
}

func (r *ArtistRepository) FindById(ctx context.Context, artistID uint64) (*models.Artist, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
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
		r.logger.Error(fmt.Sprintf("[artist repo] failed to scan columns in FindById: %v", err), requestID)
		return nil, errors.Wrap(err, "FindById.Query")
	}
	r.logger.Info("[artist repo] successful FindById query", requestID)

	return artist, nil
}

func (r *ArtistRepository) FindByName(ctx context.Context, name string) ([]*models.Artist, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var artists []*models.Artist
	rows, err := r.db.QueryContext(ctx, findByNameQuery, name)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[artist repo] failed to execute FindByName query: %v", err), requestID)
		return nil, errors.Wrap(err, "FindByName.Query")
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
			r.logger.Error(fmt.Sprintf("[artist repo] failed to scan columns in FindByName: %v", err), requestID)
			return nil, errors.Wrap(err, "FindByName.Query")
		}
		if strings.Contains(artist.Name, name) {
			artists = append(artists, artist)
		}
	}
	r.logger.Info("[artist repo] successful FindByName query", requestID)

	return artists, nil
}

func (r *ArtistRepository) GetAll(ctx context.Context) ([]*models.Artist, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var artists []*models.Artist
	rows, err := r.db.QueryContext(ctx, getAllQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[artist repo] failed to execute GetAll query: %v", err), requestID)
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
			r.logger.Error(fmt.Sprintf("[artist repo] failed to scan columns in GetAll: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAll.Query")
		}
		artists = append(artists, artist)
	}
	r.logger.Info("[artist repo] successful GetAll query", requestID)

	return artists, nil
}

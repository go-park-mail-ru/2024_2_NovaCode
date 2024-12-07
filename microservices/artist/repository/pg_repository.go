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
		r.logger.Error(fmt.Sprintf("[artist repo] failed to scan row in Create: %v", err), requestID)
		return nil, errors.Wrap(err, "Create.Scan")
	}
	r.logger.Info("[artist repo] successful Create scan row", requestID)

	return createdArtist, nil
}

func (r *ArtistRepository) FindById(ctx context.Context, artistID uint64) (*models.Artist, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	artist := &models.Artist{}

	stmt, err := r.db.PrepareContext(ctx, findByIDQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[artist repo] failed to prepare context in FindById: %v", err), requestID)
		return nil, errors.Wrap(err, "FindById.PrepareContext")
	}
	r.logger.Info("[artist repo] successful FindById prepare context", requestID)
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
		r.logger.Error(fmt.Sprintf("[artist repo] failed to scan row in FindById: %v", err), requestID)
		return nil, errors.Wrap(err, "FindById.Scan")
	}
	r.logger.Info("[artist repo] successful FindById scan row", requestID)

	return artist, nil
}

func (r *ArtistRepository) FindByQuery(ctx context.Context, query string) ([]*models.Artist, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	tsQuery := utils.MakeSearchQuery(query)

	stmt, err := r.db.PrepareContext(ctx, findByQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[artist repo] failed to prepare context in FindByQuery: %v", err), requestID)
		return nil, errors.Wrap(err, "FindByQuery.PrepareContext")
	}
	r.logger.Info("[artist repo] successful FindByQuery prepare context", requestID)
	defer stmt.Close()

	var artists []*models.Artist
	rows, err := stmt.QueryContext(ctx, tsQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[artist repo] failed to query context in FindByQuery: %v", err), requestID)
		return nil, errors.Wrap(err, "FindByQuery.Query")
	}
	r.logger.Info("[artist repo] successful FindByQuery query context", requestID)
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
			r.logger.Error(fmt.Sprintf("[artist repo] failed to scan rows in FindByQuery: %v", err), requestID)
			return nil, errors.Wrap(err, "FindByQuery.Scan")
		}
		r.logger.Info("[artist repo] successful FindByQuery scan rows", requestID)

		artists = append(artists, artist)
	}

	return artists, nil
}

func (r *ArtistRepository) GetAll(ctx context.Context) ([]*models.Artist, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var artists []*models.Artist
	rows, err := r.db.QueryContext(ctx, getAllQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[artist repo] failed to query context in GetAll: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAll.Query")
	}
	r.logger.Info("[artist repo] successful GetAll query context", requestID)
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
			r.logger.Error(fmt.Sprintf("[artist repo] failed to scan rows in GetAll: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAll.Scan")
		}
		r.logger.Info("[artist repo] successful GetAll scan rows", requestID)

		artists = append(artists, artist)
	}

	return artists, nil
}

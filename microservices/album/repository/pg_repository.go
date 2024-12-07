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

type AlbumRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewAlbumPGRepository(db *sql.DB, logger logger.Logger) *AlbumRepository {
	return &AlbumRepository{db, logger}
}

func (r *AlbumRepository) Create(ctx context.Context, album *models.Album) (*models.Album, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	createdAlbum := &models.Album{}
	row := r.db.QueryRowContext(
		ctx,
		createAlbumQuery,
		album.Name,
		album.ReleaseDate,
		album.Image,
		album.ArtistID,
	)

	if err := row.Scan(
		&createdAlbum.ID,
		&createdAlbum.Name,

		&createdAlbum.ReleaseDate,
		&createdAlbum.Image,
		&createdAlbum.ArtistID,
		&createdAlbum.CreatedAt,
		&createdAlbum.UpdatedAt,
	); err != nil {
		r.logger.Error(fmt.Sprintf("[album repo] failed to scan row in Create: %v", err), requestID)
		return nil, errors.Wrap(err, "Create.Scan")
	}
	r.logger.Info("[album repo] successful scan row", requestID)

	return createdAlbum, nil
}

func (r *AlbumRepository) FindById(ctx context.Context, albumID uint64) (*models.Album, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	stmt, err := r.db.PrepareContext(ctx, findByIDQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[album repo] failed to prepare context in FindById: %v", err), requestID)
		return nil, errors.Wrap(err, "FindById.Prepare")
	}
	r.logger.Info("[album repo] successful FindById prepare context", requestID)
	defer stmt.Close()

	album := &models.Album{}
	if err := stmt.QueryRowContext(ctx, albumID).Scan(
		&album.ID,
		&album.Name,
		&album.ReleaseDate,
		&album.Image,
		&album.ArtistID,
		&album.CreatedAt,
		&album.UpdatedAt,
	); err != nil {
		r.logger.Error(fmt.Sprintf("[album repo] failed to scan row in FindById: %v", err), requestID)
		return nil, errors.Wrap(err, "FindById.Scan")
	}
	r.logger.Info("[album repo] successful FindById scan row", requestID)

	return album, nil
}

func (r *AlbumRepository) FindByQuery(ctx context.Context, query string) ([]*models.Album, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	tsQuery := utils.MakeSearchQuery(query)

	stmt, err := r.db.PrepareContext(ctx, findByQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[album repo] failed to prepare context in FindByQuery: %v", err), requestID)
		return nil, errors.Wrap(err, "FindByQuery.Prepare")
	}
	r.logger.Info("[album repo] successful FindByQuery prepare context", requestID)
	defer stmt.Close()

	var albums []*models.Album
	rows, err := stmt.QueryContext(ctx, tsQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[album repo] failed to query in FindByQuery: %v", err), requestID)
		return nil, errors.Wrap(err, "FindByQuery.Query")
	}
	r.logger.Info("[album repo] successful FindByQuery query", requestID)
	defer rows.Close()

	for rows.Next() {
		album := &models.Album{}
		err := rows.Scan(
			&album.ID,
			&album.Name,
			&album.ReleaseDate,
			&album.Image,
			&album.ArtistID,
			&album.CreatedAt,
			&album.UpdatedAt,
		)
		if err != nil {
			r.logger.Error(fmt.Sprintf("[album repo] failed to scan rows in FindByQuery: %v", err), requestID)
			return nil, errors.Wrap(err, "FindByQuery.Scan")
		}
		r.logger.Info("[album repo] successful FindByQuery scan rows", requestID)

		albums = append(albums, album)
	}

	return albums, nil
}

func (r *AlbumRepository) GetAll(ctx context.Context) ([]*models.Album, error) {
	requestID := ctx.Value(utils.RequestIDKey{})

	var albums []*models.Album
	rows, err := r.db.QueryContext(ctx, getAllQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[album repo] failed to query context in GetAll: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAll.QueryContext")
	}
	r.logger.Info("[album repo] successful GetAll query context", requestID)
	defer rows.Close()

	for rows.Next() {
		album := &models.Album{}
		err := rows.Scan(
			&album.ID,
			&album.Name,
			&album.ReleaseDate,
			&album.Image,
			&album.ArtistID,
			&album.CreatedAt,
			&album.UpdatedAt,
		)
		if err != nil {
			r.logger.Error(fmt.Sprintf("[album repo] failed to scan rows in GetAll: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAll.Query")
		}
		r.logger.Info("[album repo] successful GetAll scan rows", requestID)
		albums = append(albums, album)
	}

	return albums, nil
}

func (r *AlbumRepository) GetAllByArtistID(ctx context.Context, artistID uint64) ([]*models.Album, error) {
	requestID := ctx.Value(utils.RequestIDKey{})

	stmt, err := r.db.PrepareContext(ctx, getByArtistIDQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[album repo] failed to prepare context in GetAllByArtistID: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAllByArtistID.Prepare")
	}
	r.logger.Info("[album repo] successful GetAllByArtistID prepare context", requestID)
	defer stmt.Close()

	var albums []*models.Album
	rows, err := stmt.QueryContext(ctx, artistID)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[album repo] failed to query context in GetAllByArtistID: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAllByArtistID.Query")
	}
	r.logger.Info("[album repo] successful GetAllByArtistID query context", requestID)
	defer rows.Close()

	for rows.Next() {
		album := &models.Album{}
		err := rows.Scan(
			&album.ID,
			&album.Name,
			&album.ReleaseDate,
			&album.Image,
			&album.ArtistID,
			&album.CreatedAt,
			&album.UpdatedAt,
		)
		if err != nil {
			r.logger.Error(fmt.Sprintf("[album repo] failed to scan rows in GetAllByArtistID: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAllByArtistID.Scan")
		}
		r.logger.Info("[album repo] successful GetAllByArtistID scan rows", requestID)
		albums = append(albums, album)
	}

	return albums, nil
}

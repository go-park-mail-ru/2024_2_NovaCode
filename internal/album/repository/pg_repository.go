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
		album.TrackCount,
		album.ReleaseDate,
		album.Image,
		album.ArtistID,
	)

	if err := row.Scan(
		&createdAlbum.ID,
		&createdAlbum.Name,
		&createdAlbum.TrackCount,
		&createdAlbum.ReleaseDate,
		&createdAlbum.Image,
		&createdAlbum.ArtistID,
		&createdAlbum.CreatedAt,
		&createdAlbum.UpdatedAt,
	); err != nil {
		r.logger.Error(fmt.Sprintf("[album repo] failed to scan columns in FindByName: %v", err), requestID)
		return nil, errors.Wrap(err, "Create.Query")
	}
	r.logger.Info("[album repo] successful Create query", requestID)

	return createdAlbum, nil
}

func (r *AlbumRepository) FindById(ctx context.Context, albumID uint64) (*models.Album, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	albums := &models.Album{}
	row := r.db.QueryRowContext(ctx, findByIDQuery, albumID)
	if err := row.Scan(
		&albums.ID,
		&albums.Name,
		&albums.TrackCount,
		&albums.ReleaseDate,
		&albums.Image,
		&albums.ArtistID,
		&albums.CreatedAt,
		&albums.UpdatedAt,
	); err != nil {
		r.logger.Error(fmt.Sprintf("[album repo] failed to scan columns in FindById: %v", err), requestID)
		return nil, errors.Wrap(err, "FindById.Query")
	}
	r.logger.Info("[album repo] successful FindById query", requestID)

	return albums, nil
}

func (r *AlbumRepository) FindByName(ctx context.Context, name string) ([]*models.Album, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var albums []*models.Album
	rows, err := r.db.QueryContext(ctx, findByNameQuery, name)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[album repo] failed to execute FindByName query: %v", err), requestID)
		return nil, errors.Wrap(err, "FindByName.Query")
	}
	defer rows.Close()

	for rows.Next() {
		album := &models.Album{}
		err := rows.Scan(
			&album.ID,
			&album.Name,
			&album.TrackCount,
			&album.ReleaseDate,
			&album.Image,
			&album.ArtistID,
			&album.CreatedAt,
			&album.UpdatedAt,
		)
		if err != nil {
			r.logger.Error(fmt.Sprintf("[repo] failed to scan columns in FindByName: %v", err), requestID)
			return nil, errors.Wrap(err, "FindByName.Query")
		}
		if strings.Contains(album.Name, name) {
			albums = append(albums, album)
		}
	}
	r.logger.Info("[album repo] successful FindByName query", requestID)

	return albums, nil
}

func (r *AlbumRepository) GetAll(ctx context.Context) ([]*models.Album, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var albums []*models.Album
	rows, err := r.db.QueryContext(ctx, getAllQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[album repo] failed to execute GetAll query: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAll.Query")
	}
	defer rows.Close()

	for rows.Next() {
		album := &models.Album{}
		err := rows.Scan(
			&album.ID,
			&album.Name,
			&album.TrackCount,
			&album.ReleaseDate,
			&album.Image,
			&album.ArtistID,
			&album.CreatedAt,
			&album.UpdatedAt,
		)
		if err != nil {
			r.logger.Error(fmt.Sprintf("[album repo] failed to scan columns in GetAll: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAll.Query")
		}
		albums = append(albums, album)
	}
	r.logger.Info("[album repo] successful GetAll query", requestID)

	return albums, nil
}

func (r *AlbumRepository) GetAllByArtistID(ctx context.Context, artistID uint64) ([]*models.Album, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var albums []*models.Album
	rows, err := r.db.QueryContext(ctx, getByArtistIDQuery, artistID)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[album repo] failed to execute GetAllByArtistID query: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAllByArtistID.Query")
	}
	defer rows.Close()

	for rows.Next() {
		album := &models.Album{}
		err := rows.Scan(
			&album.ID,
			&album.Name,
			&album.TrackCount,
			&album.ReleaseDate,
			&album.Image,
			&album.ArtistID,
			&album.CreatedAt,
			&album.UpdatedAt,
		)
		if err != nil {
			r.logger.Error(fmt.Sprintf("[album repo] failed to scan columns in GetAllByArtistID: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAllByArtistID.Query")
		}
		albums = append(albums, album)
	}
	r.logger.Info("[album repo] successful GetAllByArtistID query", requestID)

	return albums, nil
}

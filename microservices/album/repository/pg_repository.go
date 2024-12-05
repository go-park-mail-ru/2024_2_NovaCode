package repository

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/pkg/errors"
)

type AlbumRepository struct {
	db *sql.DB
}

func NewAlbumPGRepository(db *sql.DB) *AlbumRepository {
	return &AlbumRepository{db: db}
}

func (r *AlbumRepository) Create(ctx context.Context, album *models.Album) (*models.Album, error) {
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
		return nil, errors.Wrap(err, "Create.Query")
	}

	return createdAlbum, nil
}

func (r *AlbumRepository) FindById(ctx context.Context, albumID uint64) (*models.Album, error) {
	stmt, err := r.db.PrepareContext(ctx, findByIDQuery)
	if err != nil {
		return nil, errors.Wrap(err, "FindById.Prepare")
	}
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
		return nil, errors.Wrap(err, "FindById.Query")
	}

	return album, nil
}

func (r *AlbumRepository) FindByQuery(ctx context.Context, query string) ([]*models.Album, error) {
	tsQuery := utils.MakeSearchQuery(query)

	stmt, err := r.db.PrepareContext(ctx, findByQuery)
	if err != nil {
		return nil, errors.Wrap(err, "FindByQuery.Prepare")
	}
	defer stmt.Close()

	var albums []*models.Album
	rows, err := stmt.QueryContext(ctx, tsQuery)
	if err != nil {
		return nil, errors.Wrap(err, "FindByQuery.Query")
	}
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
			return nil, errors.Wrap(err, "FindByQuery.Scan")
		}

		albums = append(albums, album)
	}

	return albums, nil
}

func (r *AlbumRepository) GetAll(ctx context.Context) ([]*models.Album, error) {
	var albums []*models.Album
	rows, err := r.db.QueryContext(ctx, getAllQuery)
	if err != nil {
		return nil, errors.Wrap(err, "GetAll.Query")
	}
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
			return nil, errors.Wrap(err, "GetAll.Query")
		}
		albums = append(albums, album)
	}

	return albums, nil
}

func (r *AlbumRepository) GetAllByArtistID(ctx context.Context, artistID uint64) ([]*models.Album, error) {
	stmt, err := r.db.PrepareContext(ctx, getByArtistIDQuery)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllByArtistID.Prepare")
	}
	defer stmt.Close()

	var albums []*models.Album
	rows, err := stmt.QueryContext(ctx, artistID)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllByArtistID.Query")
	}
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
			return nil, errors.Wrap(err, "GetAllByArtistID.Scan")
		}
		albums = append(albums, album)
	}

	return albums, nil
}

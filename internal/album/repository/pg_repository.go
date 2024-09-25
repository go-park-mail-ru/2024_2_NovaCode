package album_repository

import (
	"context"
	"database/sql"
	"strings"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
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
		album.Genre,
		album.TrackCount,
		album.ReleaseDate,
		album.Image,
		album.ArtistID,
	)

	if err := row.Scan(
		&createdAlbum.ID,
		&createdAlbum.Name,
		&createdAlbum.Genre,
		&createdAlbum.TrackCount,
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
	albums := &models.Album{}
	row := r.db.QueryRowContext(ctx, findByIDQuery, albumID)
	if err := row.Scan(
		&albums.ID,
		&albums.Name,
		&albums.Genre,
		&albums.TrackCount,
		&albums.ReleaseDate,
		&albums.Image,
		&albums.ArtistID,
		&albums.CreatedAt,
		&albums.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "FindById.Query")
	}

	return albums, nil
}

func (r *AlbumRepository) FindByName(ctx context.Context, name string) ([]*models.Album, error) {
	var albums []*models.Album
	rows, err := r.db.QueryContext(ctx, findByNameQuery, name)
	if err != nil {
		return nil, errors.Wrap(err, "FindByName.Query")
	}
	defer rows.Close()

	for rows.Next() {
		album := &models.Album{}
		err := rows.Scan(
			&album.ID,
			&album.Name,
			&album.Genre,
			&album.TrackCount,
			&album.ReleaseDate,
			&album.Image,
			&album.ArtistID,
			&album.CreatedAt,
			&album.UpdatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "FindByName.Query")
		}
		if strings.Contains(album.Name, name) {
			albums = append(albums, album)
		}
	}

	return albums, nil
}

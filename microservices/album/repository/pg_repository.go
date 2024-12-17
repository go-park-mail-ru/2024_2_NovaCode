package repository

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	uuid "github.com/google/uuid"
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
	albums := &models.Album{}
	row := r.db.QueryRowContext(ctx, findByIDQuery, albumID)
	if err := row.Scan(
		&albums.ID,
		&albums.Name,
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

func (r *AlbumRepository) FindByQuery(ctx context.Context, query string) ([]*models.Album, error) {
	tsQuery := utils.MakeSearchQuery(query)

	var albums []*models.Album
	rows, err := r.db.QueryContext(ctx, findByQuery, tsQuery)
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
			return nil, errors.Wrap(err, "FindByQuery.Query")
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
	var albums []*models.Album
	rows, err := r.db.QueryContext(ctx, getByArtistIDQuery, artistID)
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
			return nil, errors.Wrap(err, "GetAllByArtistID.Query")
		}
		albums = append(albums, album)
	}

	return albums, nil
}

func (r *AlbumRepository) AddFavoriteAlbum(ctx context.Context, userID uuid.UUID, albumID uint64) error {
	_, err := r.db.ExecContext(ctx, addFavoriteAlbumQuery, userID, albumID)
	if err != nil {
		return errors.Wrap(err, "AddFavoriteAlbum.Query")
	}

	return nil
}

func (r *AlbumRepository) DeleteFavoriteAlbum(ctx context.Context, userID uuid.UUID, albumID uint64) error {
	_, err := r.db.ExecContext(ctx, deleteFavoriteAlbumQuery, userID, albumID)
	if err != nil {
		return errors.Wrap(err, "DeleteFavoriteAlbum.Query")
	}

	return nil
}

func (r *AlbumRepository) IsFavoriteAlbum(ctx context.Context, userID uuid.UUID, albumID uint64) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx, isFavoriteAlbumQuery, userID, albumID).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.Wrap(err, "IsFavoriteAlbum.Query")
	}

	return exists, nil
}

func (r *AlbumRepository) GetFavoriteAlbums(ctx context.Context, userID uuid.UUID) ([]*models.Album, error) {
	var albums []*models.Album
	rows, err := r.db.QueryContext(ctx, getFavoriteQuery, userID)
	if err != nil {
		return nil, errors.Wrap(err, "GetFavoriteAlbums.Query")
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
			return nil, errors.Wrap(err, "GetFavoriteAlbums.Query")
		}
		albums = append(albums, album)
	}

	return albums, nil
}

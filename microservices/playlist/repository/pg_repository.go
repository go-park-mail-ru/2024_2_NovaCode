package repository

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type PlaylistRepository struct {
	db *sql.DB
}

func NewPlaylistRepository(db *sql.DB) playlist.Repository {
	return &PlaylistRepository{db: db}
}

func (r *PlaylistRepository) CreatePlaylist(ctx context.Context, playlist *models.Playlist) (*models.Playlist, error) {
	insertedPlaylist := &models.Playlist{}
	row := r.db.QueryRowContext(ctx,
		CreatePlaylistQuery,
		playlist.Name,
		playlist.Image,
		playlist.OwnerID,
	)

	if err := row.Scan(
		&insertedPlaylist.ID,
		&insertedPlaylist.Name,
		&insertedPlaylist.Image,
		&insertedPlaylist.OwnerID,
		&insertedPlaylist.IsPrivate,
		&insertedPlaylist.CreatedAt,
		&insertedPlaylist.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return insertedPlaylist, nil
}

func (r *PlaylistRepository) GetAllPlaylists(ctx context.Context) ([]*models.Playlist, error) {
	playlists := []*models.Playlist{}
	rows, err := r.db.QueryContext(ctx,
		GetAllPlaylistsQuery,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		playlist := &models.Playlist{}
		if err := rows.Scan(
			&playlist.ID,
			&playlist.Name,
			&playlist.Image,
			&playlist.OwnerID,
			&playlist.IsPrivate,
			&playlist.CreatedAt,
			&playlist.UpdatedAt,
		); err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

func (r *PlaylistRepository) GetPlaylist(ctx context.Context, playlistID uint64) (*models.Playlist, error) {
	playlist := &models.Playlist{}
	row := r.db.QueryRowContext(ctx,
		GetPlaylistQuery,
		playlistID,
	)

	if err := row.Scan(
		&playlist.ID,
		&playlist.Name,
		&playlist.Image,
		&playlist.OwnerID,
		&playlist.IsPrivate,
		&playlist.CreatedAt,
		&playlist.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return playlist, nil
}

func (r *PlaylistRepository) GetLengthPlaylist(ctx context.Context, playlistID uint64) (uint64, error) {
	var length uint64
	row := r.db.QueryRowContext(ctx,
		GetLengthPlaylistsQuery,
		playlistID,
	)

	if err := row.Scan(
		&length,
	); err != nil {
		return 0, err
	}

	return length, nil
}

func (r *PlaylistRepository) GetUserPlaylists(ctx context.Context, userID uuid.UUID) ([]*models.Playlist, error) {
	playlists := []*models.Playlist{}
	rows, err := r.db.QueryContext(ctx,
		GetUserPlaylistsQuery,
		userID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		playlist := &models.Playlist{}
		if err := rows.Scan(
			&playlist.ID,
			&playlist.Name,
			&playlist.Image,
			&playlist.OwnerID,
			&playlist.IsPrivate,
			&playlist.CreatedAt,
			&playlist.UpdatedAt,
		); err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

func (r *PlaylistRepository) AddToPlaylist(ctx context.Context, playlistID uint64, trackOrder uint64, trackID uint64) (*models.PlaylistTrack, error) {
	insertedTrack := &models.PlaylistTrack{}
	row := r.db.QueryRowContext(ctx,
		AddToPlaylistQuery,
		playlistID,
		trackOrder,
		trackID,
	)

	if err := row.Scan(
		&insertedTrack.ID,
		&insertedTrack.PlaylistID,
		&insertedTrack.TrackOrderInPlaylist,
		&insertedTrack.TrackID,
		&insertedTrack.CreatedAt,
	); err != nil {
		return nil, err
	}
	return insertedTrack, nil
}

func (r *PlaylistRepository) RemoveFromPlaylist(ctx context.Context, playlistID uint64, trackID uint64) (sql.Result, error) {
	res, err := r.db.ExecContext(ctx,
		RemoveFromPlaylistQuery,
		playlistID,
		trackID,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *PlaylistRepository) DeletePlaylist(ctx context.Context, playlistID uint64) (sql.Result, error) {
	res, err := r.db.ExecContext(ctx,
		DeletePlaylistQuery,
		playlistID,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *PlaylistRepository) AddFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) error {
	_, err := r.db.ExecContext(ctx, addFavoritePlaylistQuery, userID, playlistID)
	if err != nil {
		return errors.Wrap(err, "AddFavoritePlaylist.Query")
	}

	return nil
}

func (r *PlaylistRepository) DeleteFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) error {
	_, err := r.db.ExecContext(ctx, deleteFavoritePlaylistQuery, userID, playlistID)
	if err != nil {
		return errors.Wrap(err, "DeleteFavoritePlaylist.Query")
	}

	return nil
}

func (r *PlaylistRepository) IsFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx, isFavoritePlaylistQuery, userID, playlistID).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.Wrap(err, "IsFavoritePlaylist.Query")
	}

	return exists, nil
}

func (r *PlaylistRepository) GetFavoritePlaylists(ctx context.Context, userID uuid.UUID) ([]*models.Playlist, error) {
	var playlists []*models.Playlist
	rows, err := r.db.QueryContext(ctx, getFavoriteQuery, userID)
	if err != nil {
		return nil, errors.Wrap(err, "GetFavoritePlaylists.Query")
	}
	defer rows.Close()

	for rows.Next() {
		playlist := &models.Playlist{}
		err := rows.Scan(
			&playlist.ID,
			&playlist.Name,
			&playlist.Image,
			&playlist.OwnerID,
			&playlist.IsPrivate,
			&playlist.CreatedAt,
			&playlist.UpdatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "GetFavoritePlaylists.Query")
		}
		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

func (r *PlaylistRepository) GetPopularPlaylists(ctx context.Context) ([]*models.Playlist, error) {
	playlists := []*models.Playlist{}
	rows, err := r.db.QueryContext(ctx,
		getPopularPlaylistsQuery,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		playlist := &models.Playlist{}
		if err := rows.Scan(
			&playlist.ID,
			&playlist.Name,
			&playlist.Image,
			&playlist.OwnerID,
			&playlist.IsPrivate,
			&playlist.CreatedAt,
			&playlist.UpdatedAt,
		); err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

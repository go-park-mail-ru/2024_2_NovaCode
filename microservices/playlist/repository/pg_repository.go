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

	stmt, err := r.db.PrepareContext(ctx, GetPlaylistQuery)
	if err != nil {
		return nil, errors.Wrap(err, "GetPlaylist.PrepareContext")
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, playlistID)
	if err := row.Scan(
		&playlist.ID,
		&playlist.Name,
		&playlist.Image,
		&playlist.OwnerID,
		&playlist.IsPrivate,
		&playlist.CreatedAt,
		&playlist.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "GetPlaylist.QueryRow")
	}

	return playlist, nil
}

func (r *PlaylistRepository) GetLengthPlaylist(ctx context.Context, playlistID uint64) (uint64, error) {
	stmt, err := r.db.PrepareContext(ctx, GetLengthPlaylistsQuery)
	if err != nil {
		return 0, errors.Wrap(err, "GetLengthPlaylist.PrepareContext")
	}
	defer stmt.Close()

	var length uint64
	row := stmt.QueryRowContext(ctx, playlistID)
	if err := row.Scan(&length); err != nil {
		return 0, errors.Wrap(err, "GetLengthPlaylist.QueryRow")
	}

	return length, nil
}

func (r *PlaylistRepository) GetUserPlaylists(ctx context.Context, userID uuid.UUID) ([]*models.Playlist, error) {
	stmt, err := r.db.PrepareContext(ctx, GetUserPlaylistsQuery)
	if err != nil {
		return nil, errors.Wrap(err, "GetUserPlaylists.PrepareContext")
	}
	defer stmt.Close()

	var playlists []*models.Playlist
	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "GetUserPlaylists.QueryContext")
	}
	defer rows.Close()

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
			return nil, errors.Wrap(err, "GetUserPlaylists.Scan")
		}
		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

func (r *PlaylistRepository) AddToPlaylist(ctx context.Context, playlistID uint64, trackOrder uint64, trackID uint64) (*models.PlaylistTrack, error) {
	stmt, err := r.db.PrepareContext(ctx, AddToPlaylistQuery)
	if err != nil {
		return nil, errors.Wrap(err, "AddToPlaylist.PrepareContext")
	}
	defer stmt.Close()

	insertedTrack := &models.PlaylistTrack{}
	row := stmt.QueryRowContext(ctx, playlistID, trackOrder, trackID)
	if err := row.Scan(
		&insertedTrack.ID,
		&insertedTrack.PlaylistID,
		&insertedTrack.TrackOrderInPlaylist,
		&insertedTrack.TrackID,
		&insertedTrack.CreatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "AddToPlaylist.QueryRow")
	}

	return insertedTrack, nil
}

func (r *PlaylistRepository) RemoveFromPlaylist(ctx context.Context, playlistID uint64, trackID uint64) (sql.Result, error) {
	stmt, err := r.db.PrepareContext(ctx, RemoveFromPlaylistQuery)
	if err != nil {
		return nil, errors.Wrap(err, "RemoveFromPlaylist.PrepareContext")
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, playlistID, trackID)
	if err != nil {
		return nil, errors.Wrap(err, "RemoveFromPlaylist.Exec")
	}

	return res, nil
}

func (r *PlaylistRepository) DeletePlaylist(ctx context.Context, playlistID uint64) (sql.Result, error) {
	stmt, err := r.db.PrepareContext(ctx, DeletePlaylistQuery)
	if err != nil {
		return nil, errors.Wrap(err, "DeletePlaylist.PrepareContext")
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, playlistID)
	if err != nil {
		return nil, errors.Wrap(err, "DeletePlaylist.Exec")
	}
	return res, nil
}

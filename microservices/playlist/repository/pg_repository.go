package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type PlaylistRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewPlaylistRepository(db *sql.DB, logger logger.Logger) playlist.Repository {
	return &PlaylistRepository{db, logger}
}

func (r *PlaylistRepository) CreatePlaylist(ctx context.Context, playlist *models.Playlist) (*models.Playlist, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
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
		r.logger.Error(fmt.Sprintf("[playlist repo] failed to scan row in CreatePlaylist: %v", err), requestID)
		return nil, err
	}
	r.logger.Info("[playlist repo] successful CreatePlaylist scan row", requestID)

	return insertedPlaylist, nil
}

func (r *PlaylistRepository) GetAllPlaylists(ctx context.Context) ([]*models.Playlist, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	playlists := []*models.Playlist{}
	rows, err := r.db.QueryContext(ctx,
		GetAllPlaylistsQuery,
	)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[playlist repo] failed to query context in GetAllPlaylists: %v", err), requestID)
		return nil, err
	}
	r.logger.Info("[playlist repo] successful GetAllPlaylists query context", requestID)

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
			r.logger.Error(fmt.Sprintf("[playlist repo] failed to scan rows in GetAllPlaylists: %v", err), requestID)
			return nil, err
		}
		r.logger.Info("[playlist repo] successful GetAllPlaylists scan rows", requestID)

		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

func (r *PlaylistRepository) GetPlaylist(ctx context.Context, playlistID uint64) (*models.Playlist, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	playlist := &models.Playlist{}

	stmt, err := r.db.PrepareContext(ctx, GetPlaylistQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[playlist repo] failed to prepare context in GetPlaylist: %v", err), requestID)
		return nil, errors.Wrap(err, "GetPlaylist.PrepareContext")
	}
	r.logger.Info("[playlist repo] successful GetPlaylist prepare context", requestID)
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
		r.logger.Error(fmt.Sprintf("[playlist repo] failed to scan row in GetPlaylist: %v", err), requestID)
		return nil, errors.Wrap(err, "GetPlaylist.Scan")
	}
	r.logger.Info("[playlist repo] successful GetPlaylist scan row", requestID)

	return playlist, nil
}

func (r *PlaylistRepository) GetLengthPlaylist(ctx context.Context, playlistID uint64) (uint64, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	stmt, err := r.db.PrepareContext(ctx, GetLengthPlaylistsQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[playlist repo] failed to prepare context in GetLengthPlaylist: %v", err), requestID)
		return 0, errors.Wrap(err, "GetLengthPlaylist.PrepareContext")
	}
	r.logger.Info("[playlist repo] successful GetLengthPlaylist prepare context", requestID)
	defer stmt.Close()

	var length uint64
	row := stmt.QueryRowContext(ctx, playlistID)
	if err := row.Scan(&length); err != nil {
		r.logger.Error(fmt.Sprintf("[playlist repo] failed to scan row in GetLengthPlaylist: %v", err), requestID)
		return 0, errors.Wrap(err, "GetLengthPlaylist.Scan")
	}
	r.logger.Info("[playlist repo] successful GetLengthPlaylist scan row", requestID)

	return length, nil
}

func (r *PlaylistRepository) GetUserPlaylists(ctx context.Context, userID uuid.UUID) ([]*models.Playlist, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	stmt, err := r.db.PrepareContext(ctx, GetUserPlaylistsQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[playlist repo] failed to prepare context in GetUserPlaylists: %v", err), requestID)
		return nil, errors.Wrap(err, "GetUserPlaylists.PrepareContext")
	}
	r.logger.Info("[playlist repo] successful GetUserPlaylists prepare context", requestID)
	defer stmt.Close()

	var playlists []*models.Playlist
	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[playlist repo] failed to query context in GetUserPlaylists: %v", err), requestID)
		return nil, errors.Wrap(err, "GetUserPlaylists.QueryContext")
	}
	r.logger.Info("[playlist repo] successful GetUserPlaylists query", requestID)
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
			r.logger.Error(fmt.Sprintf("[playlist repo] failed to scan columns in GetUserPlaylists: %v", err), requestID)
			return nil, errors.Wrap(err, "GetUserPlaylists.Scan")
		}
		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

func (r *PlaylistRepository) AddToPlaylist(ctx context.Context, playlistID uint64, trackOrder uint64, trackID uint64) (*models.PlaylistTrack, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	stmt, err := r.db.PrepareContext(ctx, AddToPlaylistQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[playlist repo] failed to prepare context in AddToPlaylist: %v", err), requestID)
		return nil, errors.Wrap(err, "AddToPlaylist.PrepareContext")
	}
	r.logger.Info("[playlist repo] successful AddToPlaylist prepare context", requestID)
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
		r.logger.Error(fmt.Sprintf("[playlist repo] failed to scan row in AddToPlaylist: %v", err), requestID)
		return nil, errors.Wrap(err, "AddToPlaylist.Scan")
	}
	r.logger.Info("[playlist repo] successful AddToPlaylist scan row", requestID)

	return insertedTrack, nil
}

func (r *PlaylistRepository) RemoveFromPlaylist(ctx context.Context, playlistID uint64, trackID uint64) (sql.Result, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	stmt, err := r.db.PrepareContext(ctx, RemoveFromPlaylistQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[playlist repo] failed to prepare context in RemoveFromPlaylist: %v", err), requestID)
		return nil, errors.Wrap(err, "RemoveFromPlaylist.PrepareContext")
	}
	r.logger.Info("[playlist repo] successful RemoveFromPlaylist prepare context", requestID)
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, playlistID, trackID)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[playlist repo] failed to execute in RemoveFromPlaylist: %v", err), requestID)
		return nil, errors.Wrap(err, "RemoveFromPlaylist.Exec")
	}
	r.logger.Info("[playlist repo] successful RemoveFromPlaylist exec", requestID)

	return res, nil
}

func (r *PlaylistRepository) DeletePlaylist(ctx context.Context, playlistID uint64) (sql.Result, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	stmt, err := r.db.PrepareContext(ctx, DeletePlaylistQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[playlist repo] failed to prepare context in DeletePlaylist: %v", err), requestID)
		return nil, errors.Wrap(err, "DeletePlaylist.PrepareContext")
	}
	r.logger.Info("[playlist repo] successful DeletePlaylist prepare context", requestID)
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, playlistID)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[playlist repo] failed to execute in DeletePlaylist: %v", err), requestID)
		return nil, errors.Wrap(err, "DeletePlaylist.Exec")
	}
	r.logger.Info("[playlist repo] successful DeletePlaylist exec", requestID)

	return res, nil
}

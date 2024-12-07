package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	uuid "github.com/google/uuid"
	"github.com/pkg/errors"
)

type TrackRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewTrackPGRepository(db *sql.DB, logger logger.Logger) *TrackRepository {
	return &TrackRepository{db, logger}
}

func (r *TrackRepository) Create(ctx context.Context, track *models.Track) (*models.Track, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	createdTrack := &models.Track{}
	row := r.db.QueryRowContext(
		ctx,
		createTrackQuery,
		track.Name,
		track.Duration,
		track.FilePath,
		track.Image,
		track.ArtistID,
		track.AlbumID,
		track.OrderInAlbum,
		track.ReleaseDate,
	)

	if err := row.Scan(
		&createdTrack.ID,
		&createdTrack.Name,
		&createdTrack.Duration,
		&createdTrack.FilePath,
		&createdTrack.Image,
		&createdTrack.ArtistID,
		&createdTrack.AlbumID,
		&createdTrack.OrderInAlbum,
		&createdTrack.ReleaseDate,
		&createdTrack.CreatedAt,
		&createdTrack.UpdatedAt,
	); err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to scan row in Create: %v", err), requestID)
		return nil, errors.Wrap(err, "Create.Scan")
	}
	r.logger.Info("[track repo] successful Create scan row", requestID)

	return createdTrack, nil
}

func (r *TrackRepository) FindById(ctx context.Context, trackID uint64) (*models.Track, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	stmt, err := r.db.PrepareContext(ctx, findByIDQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to prepare context in FindById: %v", err), requestID)
		return nil, errors.Wrap(err, "FindById.PrepareContext")
	}
	r.logger.Info("[track repo] successful FindById prepare context", requestID)
	defer stmt.Close()

	track := &models.Track{}
	row := stmt.QueryRowContext(ctx, trackID)
	if err := row.Scan(
		&track.ID,
		&track.Name,
		&track.Duration,
		&track.FilePath,
		&track.Image,
		&track.ArtistID,
		&track.AlbumID,
		&track.OrderInAlbum,
		&track.ReleaseDate,
		&track.CreatedAt,
		&track.UpdatedAt,
	); err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to scan row in FindById: %v", err), requestID)
		return nil, errors.Wrap(err, "FindById.Scan")
	}
	r.logger.Info("[track repo] successful FindById scan row", requestID)

	return track, nil
}

func (r *TrackRepository) FindByQuery(ctx context.Context, query string) ([]*models.Track, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	tsQuery := utils.MakeSearchQuery(query)

	stmt, err := r.db.PrepareContext(ctx, findByQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to prepare context in FindByQuery: %v", err), requestID)
		return nil, errors.Wrap(err, "FindByQuery.Prepare")
	}
	r.logger.Info("[track repo] successful FindByQuery prepare context", requestID)
	defer stmt.Close()

	var tracks []*models.Track
	rows, err := stmt.QueryContext(ctx, tsQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to query context in FindByQuery: %v", err), requestID)
		return nil, errors.Wrap(err, "FindByQuery.Query")
	}
	r.logger.Info("[track repo] successful FindByQuery query context", requestID)
	defer rows.Close()

	for rows.Next() {
		track := &models.Track{}
		err := rows.Scan(
			&track.ID,
			&track.Name,
			&track.Duration,
			&track.FilePath,
			&track.Image,
			&track.ArtistID,
			&track.AlbumID,
			&track.OrderInAlbum,
			&track.ReleaseDate,
			&track.CreatedAt,
			&track.UpdatedAt,
		)
		if err != nil {
			r.logger.Error(fmt.Sprintf("[track repo] failed to scan rows in FindByQuery: %v", err), requestID)
			return nil, errors.Wrap(err, "FindByQuery.Scan")
		}
		r.logger.Info("[track repo] successful FindByQuery scan rows", requestID)

		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (r *TrackRepository) GetAll(ctx context.Context) ([]*models.Track, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	var tracks []*models.Track
	rows, err := r.db.QueryContext(ctx, getAllQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to query context in GetAll: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAll.Query")
	}
	r.logger.Info("[track repo] successful GetAll query context", requestID)
	defer rows.Close()

	for rows.Next() {
		track := &models.Track{}
		err := rows.Scan(
			&track.ID,
			&track.Name,
			&track.Duration,
			&track.FilePath,
			&track.Image,
			&track.ArtistID,
			&track.AlbumID,
			&track.OrderInAlbum,
			&track.ReleaseDate,
			&track.CreatedAt,
			&track.UpdatedAt,
		)
		if err != nil {
			r.logger.Error(fmt.Sprintf("[track repo] failed to scan rows in GetAll: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAll.Scan")
		}
		r.logger.Info("[track repo] successful GetAll scan rows", requestID)
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (r *TrackRepository) GetAllByArtistID(ctx context.Context, artistID uint64) ([]*models.Track, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	stmt, err := r.db.PrepareContext(ctx, getByArtistIDQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to prepare context in GetAllByArtistID: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAllByArtistID.PrepareContext")
	}
	r.logger.Info("[track repo] successful GetAllByArtistID prepare context", requestID)
	defer stmt.Close()

	var tracks []*models.Track
	rows, err := stmt.QueryContext(ctx, artistID)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to query context in GetAllByArtistID: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAllByArtistID.QueryContext")
	}
	r.logger.Info("[track repo] successful GetAllByArtistID query context", requestID)
	defer rows.Close()

	for rows.Next() {
		track := &models.Track{}
		err := rows.Scan(
			&track.ID,
			&track.Name,
			&track.Duration,
			&track.FilePath,
			&track.Image,
			&track.ArtistID,
			&track.AlbumID,
			&track.OrderInAlbum,
			&track.ReleaseDate,
			&track.CreatedAt,
			&track.UpdatedAt,
		)
		if err != nil {
			r.logger.Error(fmt.Sprintf("[track repo] failed to scan rows in GetAllByArtistID: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAllByArtistID.Scan")
		}
		r.logger.Info("[track repo] successful GetAllByArtistID scan rows", requestID)

		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (r *TrackRepository) GetAllByAlbumID(ctx context.Context, albumID uint64) ([]*models.Track, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	stmt, err := r.db.PrepareContext(ctx, getByAlbumIDQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to prepare context in GetAllByAlbumID: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAllByAlbumID.PrepareContext")
	}
	r.logger.Info("[track repo] successful GetAllByAlbumID prepare context", requestID)
	defer stmt.Close()

	var tracks []*models.Track
	rows, err := stmt.QueryContext(ctx, albumID)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to query context in GetAllByAlbumID: %v", err), requestID)
		return nil, errors.Wrap(err, "GetAllByAlbumID.QueryContext")
	}
	r.logger.Info("[track repo] successful GetAllByAlbumID query context", requestID)
	defer rows.Close()

	for rows.Next() {
		track := &models.Track{}
		err := rows.Scan(
			&track.ID,
			&track.Name,
			&track.Duration,
			&track.FilePath,
			&track.Image,
			&track.ArtistID,
			&track.AlbumID,
			&track.OrderInAlbum,
			&track.ReleaseDate,
			&track.CreatedAt,
			&track.UpdatedAt,
		)
		if err != nil {
			r.logger.Error(fmt.Sprintf("[track repo] failed to scan rows in GetAllByAlbumID: %v", err), requestID)
			return nil, errors.Wrap(err, "GetAllByAlbumID.Scan")
		}
		r.logger.Info("[track repo] successful GetAllByAlbumID scan rows", requestID)

		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (r *TrackRepository) AddFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) error {
	requestID := ctx.Value(utils.RequestIDKey{})
	stmt, err := r.db.PrepareContext(ctx, addFavoriteTrackQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to prepare context in AddFavoriteTrack: %v", err), requestID)
		return errors.Wrap(err, "AddFavoriteTrack.PrepareContext")
	}
	r.logger.Info("[track repo] successful AddFavoriteTrack prepare context", requestID)
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userID, trackID)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to execute in AddFavoriteTrack: %v", err), requestID)
		return errors.Wrap(err, "AddFavoriteTrack.Exec")
	}
	r.logger.Info("[track repo] successful AddFavoriteTrack execute", requestID)

	return nil
}

func (r *TrackRepository) DeleteFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) error {
	requestID := ctx.Value(utils.RequestIDKey{})
	stmt, err := r.db.PrepareContext(ctx, deleteFavoriteTrackQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to prepare context in DeleteFavoriteTrack: %v", err), requestID)
		return errors.Wrap(err, "DeleteFavoriteTrack.PrepareContext")
	}
	r.logger.Info("[track repo] successful DeleteFavoriteTrack prepare context", requestID)
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userID, trackID)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to execute in DeleteFavoriteTrack: %v", err), requestID)
		return errors.Wrap(err, "DeleteFavoriteTrack.Exec")
	}
	r.logger.Info("[track repo] successful DeleteFavoriteTrack execute", requestID)

	return nil
}

func (r *TrackRepository) IsFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) (bool, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	stmt, err := r.db.PrepareContext(ctx, isFavoriteTrackQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to prepare context in IsFavoriteTrack: %v", err), requestID)
		return false, errors.Wrap(err, "IsFavoriteTrack.PrepareContext")
	}
	r.logger.Info("[track repo] successful IsFavoriteTrack prepare context", requestID)
	defer stmt.Close()

	var exists bool
	err = stmt.QueryRowContext(ctx, userID, trackID).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		r.logger.Error(fmt.Sprintf("[track repo] failed to scan row in IsFavoriteTrack: %v", err), requestID)
		return false, errors.Wrap(err, "IsFavoriteTrack.Scan")
	}
	r.logger.Info("[track repo] successful IsFavoriteTrack scan row", requestID)

	return exists, nil
}

func (r *TrackRepository) GetFavoriteTracks(ctx context.Context, userID uuid.UUID) ([]*models.Track, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	stmt, err := r.db.PrepareContext(ctx, getFavoriteQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to prepare context in GetFavoriteTracks: %v", err), requestID)
		return nil, errors.Wrap(err, "GetFavoriteTracks.PrepareContext")
	}
	r.logger.Info("[track repo] successful GetFavoriteTracks prepare context", requestID)
	defer stmt.Close()

	var tracks []*models.Track
	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to query context in GetFavoriteTracks: %v", err), requestID)
		return nil, errors.Wrap(err, "GetFavoriteTracks.QueryContext")
	}
	r.logger.Info("[track repo] successful GetFavoriteTracks query context", requestID)
	defer rows.Close()

	for rows.Next() {
		track := &models.Track{}
		err := rows.Scan(
			&track.ID,
			&track.Name,
			&track.Duration,
			&track.FilePath,
			&track.Image,
			&track.ArtistID,
			&track.AlbumID,
			&track.OrderInAlbum,
			&track.ReleaseDate,
			&track.CreatedAt,
			&track.UpdatedAt,
		)
		if err != nil {
			r.logger.Error(fmt.Sprintf("[track repo] failed to scan rows in GetFavoriteTracks: %v", err), requestID)
			return nil, errors.Wrap(err, "GetFavoriteTracks.Scan")
		}
		r.logger.Info("[track repo] successful GetFavoriteTracks scan rows", requestID)

		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (r *TrackRepository) GetTracksFromPlaylist(ctx context.Context, playlistID uint64) ([]*models.PlaylistTrack, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	stmt, err := r.db.PrepareContext(ctx, GetTracksFromPlaylistQuery)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to prepare context in GetTracksFromPlaylist: %v", err), requestID)
		return nil, errors.Wrap(err, "GetTracksFromPlaylist.PrepareContext")
	}
	r.logger.Info("[track repo] successful GetTracksFromPlaylist prepare context", requestID)
	defer stmt.Close()

	var playlist []*models.PlaylistTrack
	rows, err := stmt.QueryContext(ctx, playlistID)
	if err != nil {
		r.logger.Error(fmt.Sprintf("[track repo] failed to query context in GetTracksFromPlaylist: %v", err), requestID)
		return nil, errors.Wrap(err, "GetTracksFromPlaylist.QueryContext")
	}
	r.logger.Info("[track repo] successful GetTracksFromPlaylist query context", requestID)
	defer rows.Close()

	for rows.Next() {
		track := &models.PlaylistTrack{}
		if err := rows.Scan(
			&track.ID,
			&track.PlaylistID,
			&track.TrackOrderInPlaylist,
			&track.TrackID,
			&track.CreatedAt,
		); err != nil {
			r.logger.Error(fmt.Sprintf("[track repo] failed to scan rows in GetTracksFromPlaylist: %v", err), requestID)
			return nil, errors.Wrap(err, "GetTracksFromPlaylist.Scan")
		}
		r.logger.Info("[track repo] successful GetTracksFromPlaylist scan row", requestID)
		playlist = append(playlist, track)
	}

	return playlist, nil
}

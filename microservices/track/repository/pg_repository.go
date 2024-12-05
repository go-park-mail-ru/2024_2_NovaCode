package repository

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	uuid "github.com/google/uuid"
	"github.com/pkg/errors"
)

type TrackRepository struct {
	db *sql.DB
}

func NewTrackPGRepository(db *sql.DB) *TrackRepository {
	return &TrackRepository{db: db}
}

func (r *TrackRepository) Create(ctx context.Context, track *models.Track) (*models.Track, error) {
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
		return nil, errors.Wrap(err, "Create.Query")
	}

	return createdTrack, nil
}

func (r *TrackRepository) FindById(ctx context.Context, trackID uint64) (*models.Track, error) {
	stmt, err := r.db.PrepareContext(ctx, findByIDQuery)
	if err != nil {
		return nil, errors.Wrap(err, "FindById.PrepareContext")
	}
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
		return nil, errors.Wrap(err, "FindById.QueryRow")
	}

	return track, nil
}

func (r *TrackRepository) FindByQuery(ctx context.Context, query string) ([]*models.Track, error) {
	tsQuery := utils.MakeSearchQuery(query)

	stmt, err := r.db.PrepareContext(ctx, findByQuery)
	if err != nil {
		return nil, errors.Wrap(err, "FindByQuery.Prepare")
	}
	defer stmt.Close()

	var tracks []*models.Track
	rows, err := stmt.QueryContext(ctx, tsQuery)
	if err != nil {
		return nil, errors.Wrap(err, "FindByQuery.Query")
	}
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
			return nil, errors.Wrap(err, "FindByQuery.Query")
		}

		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (r *TrackRepository) GetAll(ctx context.Context) ([]*models.Track, error) {
	var tracks []*models.Track
	rows, err := r.db.QueryContext(ctx, getAllQuery)
	if err != nil {
		return nil, errors.Wrap(err, "GetAll.Query")
	}
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
			return nil, errors.Wrap(err, "GetAll.Query")
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (r *TrackRepository) GetAllByArtistID(ctx context.Context, artistID uint64) ([]*models.Track, error) {
	stmt, err := r.db.PrepareContext(ctx, getByArtistIDQuery)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllByArtistID.PrepareContext")
	}
	defer stmt.Close()

	var tracks []*models.Track
	rows, err := stmt.QueryContext(ctx, artistID)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllByArtistID.QueryContext")
	}
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
			return nil, errors.Wrap(err, "GetAllByArtistID.Scan")
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (r *TrackRepository) GetAllByAlbumID(ctx context.Context, albumID uint64) ([]*models.Track, error) {
	stmt, err := r.db.PrepareContext(ctx, getByAlbumIDQuery)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllByAlbumID.PrepareContext")
	}
	defer stmt.Close()

	var tracks []*models.Track
	rows, err := stmt.QueryContext(ctx, albumID)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllByAlbumID.QueryContext")
	}
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
			return nil, errors.Wrap(err, "GetAllByAlbumID.Scan")
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (r *TrackRepository) AddFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) error {
	stmt, err := r.db.PrepareContext(ctx, addFavoriteTrackQuery)
	if err != nil {
		return errors.Wrap(err, "AddFavoriteTrack.PrepareContext")
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userID, trackID)
	if err != nil {
		return errors.Wrap(err, "AddFavoriteTrack.Exec")
	}

	return nil
}

func (r *TrackRepository) DeleteFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) error {
	stmt, err := r.db.PrepareContext(ctx, deleteFavoriteTrackQuery)
	if err != nil {
		return errors.Wrap(err, "DeleteFavoriteTrack.PrepareContext")
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userID, trackID)
	if err != nil {
		return errors.Wrap(err, "DeleteFavoriteTrack.Exec")
	}

	return nil
}

func (r *TrackRepository) IsFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) (bool, error) {
	stmt, err := r.db.PrepareContext(ctx, isFavoriteTrackQuery)
	if err != nil {
		return false, errors.Wrap(err, "IsFavoriteTrack.PrepareContext")
	}
	defer stmt.Close()

	var exists bool
	err = stmt.QueryRowContext(ctx, userID, trackID).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.Wrap(err, "IsFavoriteTrack.QueryRow")
	}

	return exists, nil
}

func (r *TrackRepository) GetFavoriteTracks(ctx context.Context, userID uuid.UUID) ([]*models.Track, error) {
	stmt, err := r.db.PrepareContext(ctx, getFavoriteQuery)
	if err != nil {
		return nil, errors.Wrap(err, "GetFavoriteTracks.PrepareContext")
	}
	defer stmt.Close()

	var tracks []*models.Track
	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "GetFavoriteTracks.QueryContext")
	}
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
			return nil, errors.Wrap(err, "GetFavoriteTracks.Scan")
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (r *TrackRepository) GetTracksFromPlaylist(ctx context.Context, playlistID uint64) ([]*models.PlaylistTrack, error) {
	stmt, err := r.db.PrepareContext(ctx, GetTracksFromPlaylistQuery)
	if err != nil {
		return nil, errors.Wrap(err, "GetTracksFromPlaylist.PrepareContext")
	}
	defer stmt.Close()

	var playlist []*models.PlaylistTrack
	rows, err := stmt.QueryContext(ctx, playlistID)
	if err != nil {
		return nil, errors.Wrap(err, "GetTracksFromPlaylist.QueryContext")
	}
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
			return nil, errors.Wrap(err, "GetTracksFromPlaylist.Scan")
		}
		playlist = append(playlist, track)
	}

	return playlist, nil
}

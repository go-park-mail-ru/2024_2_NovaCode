package playlist

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/google/uuid"
)

type Repository interface {
	CreatePlaylist(ctx context.Context, playlist *models.Playlist) (*models.Playlist, error)
	GetAllPlaylists(ctx context.Context) ([]*models.Playlist, error)
	GetPlaylist(ctx context.Context, playlistID uint64) (*models.Playlist, error)
	GetLengthPlaylist(ctx context.Context, playlistID uint64) (uint64, error)
	GetUserPlaylists(ctx context.Context, userID uuid.UUID) ([]*models.Playlist, error)
	AddToPlaylist(ctx context.Context, playlistID uint64, trackOrder uint64, trackID uint64) (*models.PlaylistTrack, error)
	RemoveFromPlaylist(ctx context.Context, playlistID uint64, trackID uint64) (sql.Result, error)
	DeletePlaylist(ctx context.Context, playlistID uint64) (sql.Result, error)
	AddFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) error
	DeleteFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) error
	IsFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) (bool, error)
	GetFavoritePlaylists(ctx context.Context, userID uuid.UUID) ([]*models.Playlist, error)
	GetPopularPlaylists(ctx context.Context) ([]*models.Playlist, error)
}

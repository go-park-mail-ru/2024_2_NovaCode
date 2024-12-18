package playlist

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	pldto "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist/dto"
	"github.com/google/uuid"
)

type Usecase interface {
	CreatePlaylist(ctx context.Context, newPlaylistDTO *pldto.PlaylistDTO) (*pldto.PlaylistDTO, error)
	GetAllPlaylists(ctx context.Context) ([]*pldto.PlaylistDTO, error)
	GetPlaylist(ctx context.Context, playlistID uint64) (*pldto.PlaylistDTO, error)
	GetUserPlaylists(ctx context.Context, userID uuid.UUID) ([]*pldto.PlaylistDTO, error)
	AddToPlaylist(ctx context.Context, playlistTrackDTO *pldto.PlaylistTrackDTO) (*models.PlaylistTrack, error)
	RemoveFromPlaylist(ctx context.Context, playlistTrackDTO *pldto.PlaylistTrackDTO) error
	DeletePlaylist(ctx context.Context, playlistID uint64) error
	AddFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) error
	DeleteFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) error
	IsFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) (bool, error)
	GetFavoritePlaylists(ctx context.Context, userID uuid.UUID) ([]*pldto.PlaylistDTO, error)
	GetFavoritePlaylistsCount(ctx context.Context, userID uuid.UUID) (uint64, error)
	GetPlaylistLikesCount(ctx context.Context, playlistID uint64) (uint64, error)
	GetPopularPlaylists(ctx context.Context) ([]*pldto.PlaylistDTO, error)
}

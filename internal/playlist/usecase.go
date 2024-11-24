package playlist

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	pldto "github.com/go-park-mail-ru/2024_2_NovaCode/internal/playlist/dto"
	tdto "github.com/go-park-mail-ru/2024_2_NovaCode/internal/track/dto"
	"github.com/google/uuid"
)

type Usecase interface {
	CreatePlaylist(ctx context.Context, newPlaylistDTO *pldto.PlaylistDTO) (*pldto.PlaylistDTO, error)
	GetAllPlaylists(ctx context.Context) ([]*pldto.PlaylistDTO, error)
	GetPlaylist(ctx context.Context, playlistID uint64) (*pldto.PlaylistDTO, error)
	GetTracksFromPlaylist(ctx context.Context, playlistID uint64) ([]*tdto.TrackDTO, error)
	GetUserPlaylists(ctx context.Context, userID uuid.UUID) ([]*pldto.PlaylistDTO, error)
	AddToPlaylist(ctx context.Context, playlistTrackDTO *pldto.PlaylistTrackDTO) (*models.PlaylistTrack, error)
	RemoveFromPlaylist(ctx context.Context, playlistTrackDTO *pldto.PlaylistTrackDTO) error
	DeletePlaylist(ctx context.Context, playlistID uint64) error
}

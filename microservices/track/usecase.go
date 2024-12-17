package track

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track/dto"
	uuid "github.com/google/uuid"
)

type Usecase interface {
	View(ctx context.Context, trackID uint64) (*dto.TrackDTO, error)
	Search(ctx context.Context, query string) ([]*dto.TrackDTO, error)
	GetAll(ctx context.Context) ([]*dto.TrackDTO, error)
	GetAllByArtistID(ctx context.Context, artistID uint64) ([]*dto.TrackDTO, error)
	GetAllByAlbumID(ctx context.Context, albumID uint64) ([]*dto.TrackDTO, error)
	AddFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) error
	DeleteFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) error
	IsFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) (bool, error)
	GetFavoriteTracks(ctx context.Context, userID uuid.UUID) ([]*dto.TrackDTO, error)
	ConvertTrackToDTO(ctx context.Context, track *models.Track) (*dto.TrackDTO, error)
	GetTracksFromPlaylist(ctx context.Context, playlistID uint64) ([]*dto.TrackDTO, error)
	GetPopular(ctx context.Context) ([]*dto.TrackDTO, error)
}

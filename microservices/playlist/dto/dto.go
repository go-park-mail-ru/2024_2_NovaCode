package dto

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/google/uuid"
)

type PlaylistDTO struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"name"`
	Image     string    `json:"image,omitempty"`
	OwnerID   uuid.UUID `json:"owner_id,omitempty"`
	OwnerName string    `json:"owner_name,omitempty"`
}

type PlaylistTrackDTO struct {
	PlaylistID uint64 `json:"playlist_id"`
	TrackID    uint64 `json:"track_id"`
}

func NewPlaylistFromPlaylistDTO(dto *PlaylistDTO) *models.Playlist {
	return &models.Playlist{Name: dto.Name, Image: dto.Image, OwnerID: dto.OwnerID}
}

func NewPlaylistToPlaylistDTO(playlist *models.Playlist) *PlaylistDTO {
	return &PlaylistDTO{Id: playlist.ID, Name: playlist.Name, Image: playlist.Image, OwnerID: playlist.OwnerID}
}

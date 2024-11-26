package dto

import (
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

type TrackDTO struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Duration    uint64    `json:"duration"`
	FilePath    string    `json:"filepath"`
	Image       string    `json:"image"`
	ArtistName  string    `json:"artistName"`
	ArtistID    uint64    `json:"artistID"`
	AlbumName   string    `json:"albumName"`
	AlbumID     uint64    `json:"albumID"`
	ReleaseDate time.Time `json:"release"`
}

func NewTrackDTO(track *models.Track) *TrackDTO {
	return &TrackDTO{
		track.ID,
		track.Name,
		track.Duration,
		track.FilePath,
		track.Image,
		track.ReleaseDate,
	}
}

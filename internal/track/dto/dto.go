package dto

import (
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

type TrackDTO struct {
	Name        string    `json:"name"`
	Duration    uint64    `json:"duration"`
	FilePath    string    `json:"filepath"`
	Image       string    `json:"image"`
	Artist      string    `json:"artist"`
	Album       string    `json:"album"`
	ReleaseDate time.Time `json:"release"`
}

func NewTrackDTO(track *models.Track, artist *models.Artist, album *models.Album) *TrackDTO {
	return &TrackDTO{
		track.Name,
		track.Duration,
		track.FilePath,
		track.Image,
		artist.Name,
		album.Name,
		track.ReleaseDate,
	}
}

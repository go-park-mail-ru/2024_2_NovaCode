package dto

import (
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

type AlbumDTO struct {
	Name        string    `json:"name"`
	Genre       string    `json:"genre"`
	TrackCount  uint64    `json:"trackCount"`
	ReleaseDate time.Time `json:"release"`
	Image       string    `json:"image"`
	Artist      string    `json:"artistId"`
}

func NewAlbumDTO(album *models.Album, artist *models.Artist) *AlbumDTO {
	return &AlbumDTO{
		album.Name,
		album.Genre,
		album.TrackCount,
		album.ReleaseDate,
		album.Image,
		artist.Name,
	}
}

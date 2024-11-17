package dto

import (
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

type AlbumDTO struct {
	Name        string    `json:"name"`
	ReleaseDate time.Time `json:"release"`
	Image       string    `json:"image"`
	Artist      string    `json:"artistName"`
}

func NewAlbumDTO(album *models.Album, artist *models.Artist) *AlbumDTO {
	return &AlbumDTO{
		album.Name,
		album.ReleaseDate,
		album.Image,
		artist.Name,
	}
}

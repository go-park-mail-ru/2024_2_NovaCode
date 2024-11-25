package dto

import (
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

type AlbumDTO struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	ReleaseDate time.Time `json:"release"`
	Image       string    `json:"image"`
	ArtistName  string    `json:"artistName"`
	ArtistID    uint64    `json:"artistID"`
}

func NewAlbumDTO(album *models.Album, artist *models.Artist) *AlbumDTO {
	return &AlbumDTO{
		album.ID,
		album.Name,
		album.ReleaseDate,
		album.Image,
		artist.Name,
		artist.ID,
	}
}

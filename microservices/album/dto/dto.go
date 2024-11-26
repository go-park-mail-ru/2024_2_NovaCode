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

func NewAlbumDTO(album *models.Album) *AlbumDTO {
	return &AlbumDTO{
		ID:          album.ID,
		Name:        album.Name,
		ReleaseDate: album.ReleaseDate,
		Image:       album.Image,
	}
}

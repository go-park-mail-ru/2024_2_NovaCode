package dto

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

type ArtistDTO struct {
	Name    string `json:"name"`
	Bio     string `json:"bio"`
	Country string `json:"country"`
	Image   string `json:"image"`
}

func NewArtistDTO(artist *models.Artist) *ArtistDTO {
	return &ArtistDTO{
		artist.Name,
		artist.Bio,
		artist.Country,
		artist.Image,
	}
}

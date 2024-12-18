package dto

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

//easyjson:json
type ArtistDTO struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Bio     string `json:"bio"`
	Country string `json:"country"`
	Image   string `json:"image"`
}

func NewArtistDTO(artist *models.Artist) *ArtistDTO {
	return &ArtistDTO{
		artist.ID,
		artist.Name,
		artist.Bio,
		artist.Country,
		artist.Image,
	}
}

//easyjson:json
type ArtistDTOs []*ArtistDTO

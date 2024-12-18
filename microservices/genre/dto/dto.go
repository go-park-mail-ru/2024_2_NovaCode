package dto

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

//easyjson:json
type GenreDTO struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	RusName string `json:"rusName"`
}

func NewGenreDTO(genre *models.Genre) *GenreDTO {
	return &GenreDTO{
		genre.ID,
		genre.Name,
		genre.RusName,
	}
}

//easyjson:json
type GenreDTOs []*GenreDTO

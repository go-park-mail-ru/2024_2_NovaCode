package models

import (
	"time"
)

type Track struct {
	ID          uint64        `json:"id"`
	Name        string        `json:"name"`
	Genre       string        `json:"genre"`
	Duration    time.Duration `json:"duration"`
	FilePath    string        `json:"filepath"`
	Image       string        `json:"image"`
	ArtistID    uint64        `json:"artistId"`
	AlbumID     uint64        `json:"albumId"`
	ReleaseDate time.Time     `json:"realease"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}

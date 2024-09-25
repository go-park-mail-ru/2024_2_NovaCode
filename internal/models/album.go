package models

import "time"

type Album struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Genre       string    `json:"genre"`
	TrackCount  uint64    `json:"trackCount"`
	ReleaseDate time.Time `json:"release"`
	Image       string    `json:"image"`
	ArtistID    uint64    `json:"artistId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

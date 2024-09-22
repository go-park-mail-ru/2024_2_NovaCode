package models

import "time"

type Album struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"title"`
	Image       string    `json:"image"`
	ArtistID    uint64    `json:"artistId"`
	ReleaseDate time.Time `json:"release"`
	Genre       string    `json:"genre"`
	TrackCount  uint64    `json:"trackCount"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

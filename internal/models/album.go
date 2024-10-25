package models

import "time"

type Album struct {
	ID          uint64
	Name        string
	TrackCount  uint64
	ReleaseDate time.Time
	Image       string
	ArtistID    uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

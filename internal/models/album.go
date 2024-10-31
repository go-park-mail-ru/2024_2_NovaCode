package models

import "time"

type Album struct {
	ID          uint64
	Name        string
	ReleaseDate time.Time
	Image       string
	ArtistID    uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

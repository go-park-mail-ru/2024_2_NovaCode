package models

import (
	"time"
)

type Track struct {
	ID          uint64
	Name        string
	Duration    uint64
	FilePath    string
	Image       string
	ArtistID    uint64
	AlbumID     uint64
	ReleaseDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

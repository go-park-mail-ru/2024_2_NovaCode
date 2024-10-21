package models

import "time"

type Playlist struct {
	ID         uint64
	Name       string
	TrackCount uint64
	Image      string
	OwnerID    uint64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

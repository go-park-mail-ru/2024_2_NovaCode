package models

import "time"

type PlaylistTrack struct {
	ID                   uint64
	PlaylistID           uint64
	TrackOrderInPlaylist uint64
	TrackID              uint64
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

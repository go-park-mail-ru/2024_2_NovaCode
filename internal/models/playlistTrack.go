package models

import "time"

//easyjson:json
type PlaylistTrack struct {
	ID                   uint64
	PlaylistID           uint64
	TrackOrderInPlaylist uint64
	TrackID              uint64
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

//easyjson:json
type PlaylistTracks []*PlaylistTrack

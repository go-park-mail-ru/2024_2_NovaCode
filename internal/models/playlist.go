package models

import (
	"time"

	"github.com/google/uuid"
)

// //easyjson:json
type Playlist struct {
	ID        uint64
	Name      string
	Image     string
	OwnerID   uuid.UUID
	IsPrivate bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//easyjson:json
type Playlists []*Playlist

package models

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

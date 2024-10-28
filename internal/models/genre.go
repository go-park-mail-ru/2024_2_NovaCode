package models

import (
	"time"
)

type Genre struct {
	ID        uint64
	Name      string
	RusName   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

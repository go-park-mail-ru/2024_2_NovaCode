package models

import "time"

type Artist struct {
	ID        uint64
	Name      string
	Bio       string
	Country   string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

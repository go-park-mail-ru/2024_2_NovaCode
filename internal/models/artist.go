package models

import "time"

type Artist struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Bio       string    `json:"bio"`
	Country   string    `json:"country"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

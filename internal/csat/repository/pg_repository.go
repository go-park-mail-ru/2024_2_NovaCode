package repository

import (
	"database/sql"
)

type CSATRepository struct {
	db *sql.DB
}

func NewCSATPGRepository(db *sql.DB) *CSATRepository {
	return &CSATRepository{db: db}
}

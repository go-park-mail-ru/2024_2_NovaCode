package models

import "github.com/google/uuid"

type CSAT struct {
	ID    uint64
	Topic string
}

type CSATQuestion struct {
	ID       uint64
	Question string
	CSATID   uint64
}

type CSATAnswer struct {
	ID             uint64
	Score          int
	UserID         uuid.UUID
	CSATQuestionID uint64
	CSATID         uint64
}

package models

import "github.com/google/uuid"

type CSATQuestion struct {
	ID       uint64
	Question string
	CSATID   uint64
}

type CSATAnswer struct {
	ID             uint64
	Score          uint8
	UserID         uuid.UUID
	CSATQuestionID uint64
}

type CSATStat struct {
	ID           uint64
	Topic        string
	QuestionID   uint64
	Question     string
	AverageScore float64
}

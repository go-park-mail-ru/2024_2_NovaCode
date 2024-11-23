package dto

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/google/uuid"
)

type CSATQuestionDTO struct {
	ID       uint64 `json:"id"`
	Question string `json:"question"`
}

func NewCSATQuestionDTO(csatQuestion *models.CSATQuestion) *CSATQuestionDTO {
	return &CSATQuestionDTO{
		csatQuestion.ID,
		csatQuestion.Question,
	}
}

type CSATAnswerDTO struct {
	ID             uint64    `json:"id,omitempty"`
	Score          uint8     `json:"score"`
	UserID         uuid.UUID `json:"user_id,omitempty"`
	CSATQuestionID uint64    `json:"question_id,omitempty"`
}

func NewCSATAnswerDTO(csatAnswer *models.CSATAnswer) *CSATAnswerDTO {
	return &CSATAnswerDTO{
		csatAnswer.ID,
		csatAnswer.Score,
		csatAnswer.UserID,
		csatAnswer.CSATQuestionID,
	}
}

func NewAnswerFromCSATAnswerDTO(answerDTO *CSATAnswerDTO) *models.CSATAnswer {
	return &models.CSATAnswer{
		Score:          answerDTO.Score,
		UserID:         answerDTO.UserID,
		CSATQuestionID: answerDTO.CSATQuestionID,
	}
}

type CSATStatisticsDTO struct {
	Topic        string
	Question     string
	AverageScore float64
}

func NewCSATStatisticsDTO(csatStat *models.CSATStat) *CSATStatisticsDTO {
	return &CSATStatisticsDTO{
		Topic:        csatStat.Topic,
		Question:     csatStat.Question,
		AverageScore: csatStat.AverageScore,
	}
}

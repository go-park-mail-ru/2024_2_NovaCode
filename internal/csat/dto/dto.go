package dto

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
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
	ID uint64 `json:"id"`
}

func NewCSATAnswerDTO(csatAnswer *models.CSATQuestion) *CSATAnswerDTO {
	return &CSATAnswerDTO{
		csatAnswer.ID,
	}
}

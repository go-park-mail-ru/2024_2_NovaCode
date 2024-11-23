package dto

import "github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"

type CSATStatisticsDTO struct {
	Topic    string                 `json:"topic"`
	Question []*CSATQuestionStatDTO `json:"questions"`
}

func NewCSATStatisticsDTO(stat *models.CSAT, questions []*CSATQuestionStatDTO) *CSATStatisticsDTO {
	return &CSATStatisticsDTO{
		stat.Topic,
		questions,
	}
}

type CSATQuestionDTO struct {
	ID       uint64 `json:"id"`
	Question string `json:"question"`
}

func NewCSATQuestionDTO(question *models.CSATQuestion) *CSATQuestionDTO {
	return &CSATQuestionDTO{
		question.ID,
		question.Question,
	}
}

type CSATQuestionStatDTO struct {
	*CSATQuestionDTO
	AverageScore float64 `json:"score"`
}

func NewCSATQuestionStatDTO(question *CSATQuestionDTO, avgScore float64) *CSATQuestionStatDTO {
	return &CSATQuestionStatDTO{
		question,
		avgScore,
	}
}

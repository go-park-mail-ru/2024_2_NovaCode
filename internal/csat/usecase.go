package csat

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/csat/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

type Usecase interface {
	GetStatistics(ctx context.Context) ([]*dto.CSATStatisticsDTO, error)
	GetQuestionsByTopic(ctx context.Context, topic string) ([]*dto.CSATQuestionDTO, error)
	SubmitAnswer(ctx context.Context, csatAnswer *models.CSATAnswer) (*dto.CSATAnswerDTO, error)
}

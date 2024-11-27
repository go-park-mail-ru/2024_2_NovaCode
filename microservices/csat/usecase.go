package csat

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/csat/dto"
)

type Usecase interface {
	GetStatistics(ctx context.Context) ([]*dto.CSATStatisticsDTO, error)
	GetQuestionsByTopic(ctx context.Context, topic string) ([]*dto.CSATQuestionDTO, error)
	SubmitAnswer(ctx context.Context, csatAnswer *models.CSATAnswer) (*dto.CSATAnswerDTO, error)
}

package csat

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
)

type Repo interface {
	GetStatistics(ctx context.Context) ([]*models.CSATStat, error)
	GetQuestionsByTopic(ctx context.Context, topic string) ([]*models.CSATQuestion, error)
	InsertAnswer(ctx context.Context, answer *models.CSATAnswer) (*models.CSATAnswer, error)
}

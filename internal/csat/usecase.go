package csat

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/csat/dto"
)

type Usecase interface {
	GetStatistics(ctx context.Context) ([]*dto.CSATStatisticsDTO, error)
}

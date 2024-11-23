package usecase

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/csat"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/csat/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

type csatUsecase struct {
	csatRepo csat.Repo
	logger   logger.Logger
}

func NewCSATUsecase(csatRepo csat.Repo, logger logger.Logger) csat.Usecase {
	return &csatUsecase{csatRepo, logger}
}

func (usecase *csatUsecase) GetStatistics(ctx context.Context) ([]*dto.CSATStatisticsDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	stats, err := usecase.csatRepo.GetStatistics(ctx)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load statistics: %v", err), requestID)
		return nil, fmt.Errorf("Can't load statistics")
	}

	var dtoStats []*dto.CSATStatisticsDTO
	for _, stat := range stats {
		dtoStat, err := usecase.convertStatisticsToDTO(stat)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for statistics: %v", err), requestID)
			return nil, fmt.Errorf("Can't create DTO for statistics")
		}
		dtoStats = append(dtoStats, dtoStat)
	}

	return dtoStats, nil
}

func (usecase *csatUsecase) convertStatisticsToDTO(stat *models.CSATStat) (*dto.CSATStatisticsDTO, error) {
	return dto.NewCSATStatisticsDTO(stat), nil
}

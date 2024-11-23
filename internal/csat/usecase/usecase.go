package usecase

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/csat"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/csat/dto"
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
		usecase.logger.Warn(fmt.Sprintf("Can't get statistics"))
	}
  
  var dtoStats []*dto.CSATStatisticsDTO
  for _, stat := range stats {
    dtoStat, err := usecase.convertStatToDTO(ctx, stat)
  }

	return dto.NewCSATStatisticsDTO()
}


func (usecase *csatUsecase) convertStatToDTO(ctx context.Context(), )


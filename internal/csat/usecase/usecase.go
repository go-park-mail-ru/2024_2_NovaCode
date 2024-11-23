package usecase

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/csat"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

type csatUsecase struct {
	csatRepo csat.Repo
	logger   logger.Logger
}

func NewCSATUsecase(csatRepo csat.Repo, logger logger.Logger) csat.Usecase {
	return &csatUsecase{csatRepo, logger}
}

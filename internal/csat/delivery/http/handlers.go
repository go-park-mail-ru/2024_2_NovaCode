package http

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/csat"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

type csatHandlers struct {
	usecase csat.Usecase
	logger  logger.Logger
}

func NewCSATHandlers(usecase csat.Usecase, logger logger.Logger) csat.Handlers {
	return &csatHandlers{usecase, logger}
}
